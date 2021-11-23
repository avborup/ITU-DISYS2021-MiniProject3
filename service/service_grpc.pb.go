// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionClient interface {
	MakeBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*Result, error)
	GetResult(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Result, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) MakeBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/service.auction/MakeBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) GetResult(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/service.auction/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility
type AuctionServer interface {
	MakeBid(context.Context, *Bid) (*Result, error)
	GetResult(context.Context, *empty.Empty) (*Result, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServer struct {
}

func (UnimplementedAuctionServer) MakeBid(context.Context, *Bid) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeBid not implemented")
}
func (UnimplementedAuctionServer) GetResult(context.Context, *empty.Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_MakeBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).MakeBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.auction/MakeBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).MakeBid(ctx, req.(*Bid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.auction/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).GetResult(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeBid",
			Handler:    _Auction_MakeBid_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _Auction_GetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
