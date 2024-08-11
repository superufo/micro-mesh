package main

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	protobuf "zhugedaojia.com/common/net/pb/user"
)

func main() {
	ctx := gctx.New()
	//add, err := gcfg.Instance().Get(ctx, "server.grpc.Address")
	name, err := gcfg.Instance().Get(ctx, "server.grpc.Name")

	var (
		conn          = grpcx.Client.MustNewGrpcClientConn(name.String())
		client        = protobuf.NewUserClient(conn)
		tp     uint32 = 1
	)

	userLoginTos := &protobuf.UserLoginTos{
		Account:  "admin",
		Password: "aa6e08f11f335af8b1b973d73ecb0ba8",
		OrgCode:  nil,
		Type:     &tp,
	}

	res, err := client.GetToken(ctx, userLoginTos)
	if err != nil {
		g.Log().Info(ctx,"err:",err)
		return
	}

	g.Log().Debug(ctx, "res.Data:", res.Data, "res.Msg:", res.Msg)
	return
}
