package multistore

import (
	"cosmossdk.io/store/v2"
	"cosmossdk.io/store/v2/commitment"
	ics23 "github.com/cosmos/ics23/go"
)

// MultiStore defines an abstraction layer containing a State Storage (SS) engine
// and one or more State Commitment (SC) engines.
//
// TODO: Move this type to the Core package.
type MultiStore interface {
	GetProof(storeKey string, version uint64, key []byte) (*ics23.CommitmentProof, error)
	LoadVersion(version uint64) error
	WorkingHash() []byte
	Commit() ([]byte, error)
}

type Store struct {
	ss store.VersionedDatabase
	sc map[string]*commitment.Database
}