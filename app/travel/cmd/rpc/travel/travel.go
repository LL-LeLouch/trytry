// Code generated by goctl. DO NOT EDIT!
// Source: travel.proto

package travel

import (
	"context"

	"trytry/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Homestay           = pb.Homestay
	HomestayDetailReq  = pb.HomestayDetailReq
	HomestayDetailResp = pb.HomestayDetailResp

	Travel interface {
		// homestayDetail
		HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error)
	}

	defaultTravel struct {
		cli zrpc.Client
	}
)

func NewTravel(cli zrpc.Client) Travel {
	return &defaultTravel{
		cli: cli,
	}
}

// homestayDetail
func (m *defaultTravel) HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error) {
	client := pb.NewTravelClient(m.cli.Conn())
	return client.HomestayDetail(ctx, in, opts...)
}