package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"time"

	"zhugedaojia.com/common/share"
	"zhugedaojia.com/gateway/server/base"
	"zhugedaojia.com/gateway/server/inf"
)

// 获取网关的服务配置
func init()  {

}

type  HttpSrv struct {
	// 数据
	Ctx context.Context

	Host    string
    Port    int
	properties HttpProperties

	base.BaseServer
}

func (h HttpSrv)Receive() []byte{
	return nil
}

func (h HttpSrv)Send([]byte) error {
	return nil
}

func (h HttpSrv)Run(ops inf.Properties) {
	h.Ctx = context.Background()
	s := g.Server()

	httpproxy := HttpToGrpcAdapt

	g.Server().Group("/backend",func(group *ghttp.RouterGroup) {
		//group.Middleware(service.Middleware.CtxInit)

		group.Group("/",


		group.Group("/", func(group *ghttp.RouterGroup) {
			// 限流熔断的中间件

			// 授权中间件
			group.Middleware(share.AuthMiddleware.HttpTokenVerify)

			group.GET("v1/base/upload", func(){
				  //服务发现 注册服务的名称默认 v1-base-upload-协议
			})
		})
	})

	g.Log().Info(h.Ctx,fmt.Sprint("ws ~~~ %s:%d",h.Host,h.Port))
	s.SetPort(h.Port)
	s.Run()
}

func (h HttpSrv) GetProperties() inf.Properties {
	return h.properties
}

func (h HttpSrv) SetProperties(properties inf.Properties){
	h.properties = properties.(HttpProperties)
}

func NewHttpSrv(host string,port int) inf.IServer {
	httpSrv := HttpSrv{
		Ctx:  gctx.New(),
		Host: host,
		Port: port,
		properties:  HttpProperties{},
	}

	httpSrv.Name = "http"
	return httpSrv
}

func  WithHttpTimeout(timeout  time.Duration) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(HttpProperties)
		hp.timeout = timeout
	}
}

func (h HttpSrv)WithHttpCash(caching  bool) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(HttpProperties)
		hp.caching = caching
	}
}

type HttpProperties struct {
	timeout time.Duration
	caching bool
}

func (h  HttpProperties)Get(key string) (val interface{}){
	return nil
}

// CORS 允许接口跨域请求
// s.BindMiddleware("/api/*", CORS)
//func CORS(r *ghttp.Request) {
//	// 使用框架默认的 CORS 设置
//	response := r.Response
//	option := response.DefaultCORSOptions()
//	option.AllowHeaders = option.AllowHeaders + ",x-applet-type"
//	response.CORS(option)
//	if r.Method == "OPTIONS" {
//		r.Response.WriteStatusExit(http.StatusOK)
//	} else {
//		r.Middleware.Next()
//	}
//}












