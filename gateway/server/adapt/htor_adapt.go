package adapt

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/ghttp"
	protobuf "zhugedaojia.com/common/net/pb/user"
	"zhugedaojia.com/common/response"

	gwPb "zhugedaojia.com/common/net/pb/gw"
	"zhugedaojia.com/common/response"
	"google.golang.org/protobuf/encoding/protojson"
)

// http转rpc协议的适配器
type HttpToGrpcAdapt struct {
	protoAdapt

	ghttp.RouterGroup

    GGC 	gwPb.ProxyClient        //grpcGwClient
}

// 代理端转发给微服务端
func (gg HttpToGrpcAdapt)Send() error{
	return nil
}

// 代理端转发给微服务端
func (gg HttpToGrpcAdapt)Reply() ([]byte,error){
	return nil,nil
}

func (gg HttpToGrpcAdapt)Run() error{
	return nil
}

func (gg *HttpToGrpcAdapt)GetRpcRes(pattern string)(response.JsonRes, error){
	resJson :=	response.JsonRes{}
	//根据pattern,  nacos服务发现获取地址


	// 根据pattern,前端的参数  获取后端grpc地址
	conn := grpcx.Client.MustNewGrpcClientConn("127.0.0.1:1111")
	gg.GGC  = gwPb.NewProxyClient(conn)
	tosPb := &gwPb.UnitaryTos{
		Msg:  0,
		Data: nil,
	}

	// 调用后端grpc 微服务
	ctx := context.Background()
	res, err := gg.GGC.ProxyUnitaryMsg(ctx, tosPb)

	err = json.Unmarshal(res.Data,&resJson)
	if err !=nil  {
		err = errors.New("编解码错误")
	}
	
	return res,err
}

func (gg *HttpToGrpcAdapt) GET(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
    resJson, err := gg.GetRpcRes(pattern)

	grg :=   ghttp.RouterGroup{}
	rg := grg.GET(pattern, func(r *ghttp.Request) {
		if err !=nil || resJson.Code!=0 {
			response.JsonError(r, err)
		}else{
			response.Json(r, response.ResCodeSuccess, "成功", res)
		}
	})

	return rg
}

// PUT registers an http handler to give the route pattern and the http method: PUT.
func (gg *HttpToGrpcAdapt) PUT(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get

	grg :=   ghttp.RouterGroup{}
	rg := grg.PUT(pattern, func() {
	})

	return rg
}

// POST registers an http handler to give the route pattern and the http method: POST.
func (gg *HttpToGrpcAdapt) POST(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.POST(pattern, func() {
	})

	return rg
}

// DELETE registers an http handler to give the route pattern and the http method: DELETE.
func (gg *HttpToGrpcAdapt) DELETE(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.DELETE(pattern, func() {

	})

	return rg
}

// PATCH registers an http handler to give the route pattern and the http method: PATCH.
func (gg *HttpToGrpcAdapt) PATCH(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.PATCH(pattern,object,params)

	return rg

}

// HEAD registers an http handler to give the route pattern and the http method: HEAD.
func (gg *HttpToGrpcAdapt) HEAD(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.HEAD(pattern,object,params)

	return rg
}

// CONNECT registers an http handler to give the route pattern and the http method: CONNECT.
func (gg *HttpToGrpcAdapt) CONNECT(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.CONNECT(pattern,object,params)

	return rg
}

// OPTIONS register an http handler to give the route pattern and the http method: OPTIONS.
func (gg *HttpToGrpcAdapt) OPTIONS(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.OPTIONS(pattern,object,params)

	return rg
}

// TRACE registers an http handler to give the route pattern and the http method: TRACE.
func (gg *HttpToGrpcAdapt) TRACE(pattern string, object interface{}, params ...interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.TRACE(pattern,object,params)

	return rg
}

// REST registers an http handler to give the route pattern according to REST rule.
func (gg *HttpToGrpcAdapt) REST(pattern string, object interface{}) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg :=    grg.REST(pattern,object)

	return rg
}

// Hook registers a hook to given route pattern.
func (gg *HttpToGrpcAdapt) Hook(pattern string, hook ghttp.HookName, handler ghttp.HandlerFunc) *ghttp.RouterGroup {
	// 根据pattern，前端的参数 调用grpc

	// 调用后端grpc 微服务

	// 将grpc 的结果传入到 Get


	grg :=   ghttp.RouterGroup{}
	rg := grg.Hook(pattern,hook ,handler)

	return rg
}

// Middleware binds one or more middleware to the router group.
func (gg *HttpToGrpcAdapt) Middleware(handlers ...ghttp.HandlerFunc) *ghttp.RouterGroup {
	grg :=   ghttp.RouterGroup{}

	return  grg.Middleware(handlers...)
}













