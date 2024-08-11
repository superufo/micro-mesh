package share

import (
	"errors"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	protobuf "zhugedaojia.com/common/net/pb/user"

	"zhugedaojia.com/common/response"
	"zhugedaojia.com/common/model"
	"zhugedaojia.com/common/service"
	"github.com/gogf/gf/v2/errors/gcode"

	gpb "zhugedaojia.com/common/net/pb/gw"
)

var AuthMiddleware = authMiddleware{}

type authMiddleware struct {
		TokenServiceUrl  string   //nacos服务发现后的jwt服务地址
}

// 上下文对象初始化
func (s *authMiddleware) CtxInit(r *ghttp.Request) {
	//业务自定义上下文
	customCtx := &model.UserContext{
		Data: make(g.Map),
	}
	// 初始化，务必最开始执行
	service.ContextService.Init(r, customCtx)

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 授权检查
func (s *authMiddleware) HttpTokenVerify(r *ghttp.Request) {
	var (
		ctx = gctx.New()

		conn          = grpcx.Client.MustNewGrpcClientConn(s.TokenServiceUrl)
		client        = protobuf.NewUserClient(conn)
	)

	//获取token
	token := r.Header.Get("authorization")
	if token!="" {
		g.Log().Debug(ctx, "res.Data:", token)

		tokenTos := &protobuf.VerifyTokenTos{Token: token}
		// grpc 调用
		res, err := client.VerifyToken(ctx, tokenTos)
		if err != nil {
			g.Log().Info(ctx, "err:", err)
			response.JsonExit(r, gcode.New(int(res.Code), err.Error(), nil), err.Error())
			return
		}

		if res.Data == false {
			response.JsonExit(r, gcode.New(int(res.Code), res.Msg, nil), res.Msg)
			return
		}

		r.Middleware.Next()
	} else {
		response.JsonExit(r, response.ResAuthError, "无权限访问")
	}
}

// 授权检查
func (s *authMiddleware) PermissionVerify(r *ghttp.Request) {
	//todo httpclient  请求权限接口

	r.Middleware.Next()
}