package base

import (
	"context"
	"fmt"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/header"
	"cosmossdk.io/x/accounts/accountstd"
	aa_interface_v1 "cosmossdk.io/x/accounts/interfaces/account_abstraction/v1"
	accountsv1 "cosmossdk.io/x/accounts/v1"
	"cosmossdk.io/x/tx/signing"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Account implements a base account.
type Account struct {
	PubKey   collections.Item[secp256k1.PubKey]
	Sequence collections.Sequence

	addrCodec address.Codec
	hs        header.Service

	signingHandlers signing.HandlerMap
}

// Authenticate implements the authentication flow of an abstracted account.
func (a Account) Authenticate(ctx context.Context, msg *aa_interface_v1.MsgAuthenticate) (*aa_interface_v1.MsgAuthenticateResponse, error) {
	if !accountstd.SenderIsAccountsModule(ctx) {
		return nil, fmt.Errorf("unauthorized: only accounts module is allowed to call this")
	}

	pubKey, signerData, err := a.computeSignerData(ctx)
	if err != nil {
		return nil, err
	}

	txData, err := a.getTxData(msg)
	if err != nil {
		return nil, err
	}

	gotSeq := msg.Tx.AuthInfo.SignerInfos[msg.SignerIndex].Sequence
	if gotSeq != signerData.Sequence {
		return nil, fmt.Errorf("unexpected sequence number, wanted: %d, got: %d", signerData.Sequence, gotSeq)
	}

	signMode, err := parseSignMode(msg.Tx.AuthInfo.SignerInfos[msg.SignerIndex].ModeInfo)
	if err != nil {
		return nil, fmt.Errorf("unable to parse sign mode: %w", err)
	}

	signature := msg.Tx.Signatures[msg.SignerIndex]

	signBytes, err := a.signingHandlers.GetSignBytes(ctx, signMode, signerData, txData)
	if err != nil {
		return nil, err
	}

	if !pubKey.VerifySignature(signBytes, signature) {
		return nil, fmt.Errorf("signature verification failed")
	}

	return &aa_interface_v1.MsgAuthenticateResponse{}, nil
}

func parseSignMode(info *tx.ModeInfo) (signingv1beta1.SignMode, error) {
	single, ok := info.Sum.(*tx.ModeInfo_Single_)
	if !ok {
		return 0, fmt.Errorf("only sign mode single accepted got: %v", info.Sum)
	}
	return signingv1beta1.SignMode(single.Single.Mode), nil
}

// computeSignerData will populate signer data and also increase the sequence.
func (a Account) computeSignerData(ctx context.Context) (secp256k1.PubKey, signing.SignerData, error) {
	addrStr, err := a.addrCodec.BytesToString(accountstd.Whoami(ctx))
	if err != nil {
		return secp256k1.PubKey{}, signing.SignerData{}, err
	}
	chainID := a.hs.GetHeaderInfo(ctx).ChainID

	wantSequence, err := a.Sequence.Next(ctx)
	if err != nil {
		return secp256k1.PubKey{}, signing.SignerData{}, err
	}

	pk, err := a.PubKey.Get(ctx)
	if err != nil {
		return secp256k1.PubKey{}, signing.SignerData{}, err
	}

	pkAny, err := codectypes.NewAnyWithValue(&pk)
	if err != nil {
		return secp256k1.PubKey{}, signing.SignerData{}, err
	}

	accNum, err := a.getNumber(ctx, addrStr)
	if err != nil {
		return secp256k1.PubKey{}, signing.SignerData{}, err
	}

	return pk, signing.SignerData{
		Address:       addrStr,
		ChainID:       chainID,
		AccountNumber: accNum,
		Sequence:      wantSequence,
		PubKey: &anypb.Any{
			TypeUrl: pkAny.TypeUrl,
			Value:   pkAny.Value,
		},
	}, nil
}

func (a Account) getNumber(ctx context.Context, addrStr string) (uint64, error) {
	accNum, err := accountstd.QueryModule[accountsv1.AccountNumberResponse](ctx, accountsv1.AccountNumberRequest{Address: addrStr})
	if err != nil {
		return 0, err
	}

	return accNum.Number, nil
}

func (a Account) getTxData(msg *aa_interface_v1.MsgAuthenticate) (signing.TxData, error) {
	// TODO: add a faster way to do this, we can avoid unmarshalling but we need
	// to write a function that converts this into the protov2 counterparty.
	txBody := new(txv1beta1.TxBody)
	err := proto.Unmarshal(msg.RawTx.BodyBytes, txBody)
	if err != nil {
		return signing.TxData{}, err
	}

	authInfo := new(txv1beta1.AuthInfo)
	err = proto.Unmarshal(msg.RawTx.AuthInfoBytes, authInfo)
	if err != nil {
		return signing.TxData{}, err
	}

	return signing.TxData{
		Body:                       txBody,
		AuthInfo:                   authInfo,
		BodyBytes:                  msg.RawTx.BodyBytes,
		AuthInfoBytes:              msg.RawTx.AuthInfoBytes,
		BodyHasUnknownNonCriticals: false, // NOTE: amino signing must be disabled.
	}, nil
}
