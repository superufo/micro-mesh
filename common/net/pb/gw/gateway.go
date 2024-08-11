package gw

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	protobuf "zhugedaojia.com/common/net/pb/gw"
)

var gctx context.Context

// var UserGrpcController = userGrpcController{}  Token(context.Context, *UserLoginTos) (*ReplyToc, error)  pb.UnimplementedUserServer pb.UnimplementedUserServer
func Register(s *grpcx.GrpcServer, ctx context.Context) {
	gctx = ctx
	protobuf.RegisterProxyServer(s.Server, &GrpcGatewayController{})
}

type GrpcGatewayController struct {
	protobuf.UnimplementedProxyServer
}

func (s GrpcGatewayController) mustEmbedUnimplementedUserServer() {

}

func (s GrpcGatewayController) ProxyUnitaryMsg(ctx context.Context,d *UnitaryTos) (*UnitaryTog, error) {
	res  :=  &UnitaryTog{
		Msg:          d.Msg,
		Data:          d.Data,
	}

	return res, nil
}

func (s GrpcGatewayController) ProxyStreamMsg(Proxy_ProxyStreamMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method ProxyStreamMsg not implemented")
}






