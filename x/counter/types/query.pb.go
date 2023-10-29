// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/counter/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryGetCountRequest defines the request type for querying x/mock count.
type QueryGetCountRequest struct {
}

func (m *QueryGetCountRequest) Reset()         { *m = QueryGetCountRequest{} }
func (m *QueryGetCountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCountRequest) ProtoMessage()    {}
func (*QueryGetCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd21727562626c9f, []int{0}
}
func (m *QueryGetCountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCountRequest.Merge(m, src)
}
func (m *QueryGetCountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCountRequest proto.InternalMessageInfo

// QueryGetCountResponse defines the response type for querying x/mock count.
type QueryGetCountResponse struct {
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (m *QueryGetCountResponse) Reset()         { *m = QueryGetCountResponse{} }
func (m *QueryGetCountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCountResponse) ProtoMessage()    {}
func (*QueryGetCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd21727562626c9f, []int{1}
}
func (m *QueryGetCountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCountResponse.Merge(m, src)
}
func (m *QueryGetCountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCountResponse proto.InternalMessageInfo

func (m *QueryGetCountResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryGetCountRequest)(nil), "cosmos.counter.v1.QueryGetCountRequest")
	proto.RegisterType((*QueryGetCountResponse)(nil), "cosmos.counter.v1.QueryGetCountResponse")
}

func init() { proto.RegisterFile("cosmos/counter/v1/query.proto", fileDescriptor_bd21727562626c9f) }

var fileDescriptor_bd21727562626c9f = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x48, 0xeb, 0x41, 0xa5,
	0xf5, 0xca, 0x0c, 0x95, 0xc4, 0xb8, 0x44, 0x02, 0x41, 0x2a, 0xdc, 0x53, 0x4b, 0x9c, 0x41, 0xa2,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0x16, 0x5c, 0xa2, 0x68, 0xe2, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x42, 0xf2, 0x5c, 0xdc, 0x25, 0xf9, 0x25, 0x89, 0x39, 0xf1, 0x60, 0x43, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xb8, 0xc0, 0x42, 0x60, 0x85, 0x46, 0x69, 0x5c, 0xac, 0x60,
	0x9d, 0x42, 0xb1, 0x5c, 0x1c, 0x30, 0xdd, 0x42, 0xea, 0x7a, 0x18, 0x56, 0xeb, 0x61, 0xb3, 0x57,
	0x4a, 0x83, 0xb0, 0x42, 0x88, 0x43, 0x9c, 0xdc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0x4a, 0x27, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x1e,
	0x20, 0x20, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0x02, 0x1e, 0x3a, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0xe0, 0xb0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x22, 0x18, 0xa2, 0xf4, 0x3c,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// GetCount queries the parameters of x/Counter module.
	GetCount(ctx context.Context, in *QueryGetCountRequest, opts ...grpc.CallOption) (*QueryGetCountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetCount(ctx context.Context, in *QueryGetCountRequest, opts ...grpc.CallOption) (*QueryGetCountResponse, error) {
	out := new(QueryGetCountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.counter.v1.Query/GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// GetCount queries the parameters of x/Counter module.
	GetCount(context.Context, *QueryGetCountRequest) (*QueryGetCountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetCount(ctx context.Context, req *QueryGetCountRequest) (*QueryGetCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.counter.v1.Query/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetCount(ctx, req.(*QueryGetCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.counter.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _Query_GetCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/counter/v1/query.proto",
}

func (m *QueryGetCountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetCountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalCount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.TotalCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetCountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalCount != 0 {
		n += 1 + sovQuery(uint64(m.TotalCount))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
			}
			m.TotalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
