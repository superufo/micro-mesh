package cmd

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"zhugedaojia.com/jwt/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"

	. "zhugedaojia.com/common/share"
)


var (
	HttpCmd = gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "start http server of jwt",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			add, err := gcfg.Instance().Get(ctx, "server.Address")
			g.Log().Info(ctx, "add:", add)
			address := add.Interfaces()

			if address==nil{
				g.Log().Error(ctx, "address为空")
			}

			addresses := strings.Split(address[0].(string), ":")
			port, _ := strconv.Atoi(addresses[1])

			//错误处理
			s.Use(ErrorMiddleware.ErrorHandler)
			s.BindMiddleware("/*", CORS)
			s.SetPort(port)

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 设置ctx 中间参数
				group.Middleware(AuthMiddleware.CtxInit)

				group.POST("/v1/user/sign-in-by-account", user.User.Login)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(AuthMiddleware.HttpTokenVerify)
					group.ALLMap(g.Map{
						"/v1/user/info": user.User.Info,
					})
				})
			})

			s.Run()
			return nil
		},
	}
)

// CORS 添加小程序允许的头 x-applet-type
func CORS(r *ghttp.Request) {
	// 使用框架默认的 CORS 设置
	response := r.Response
	option := response.DefaultCORSOptions()
	option.AllowHeaders = option.AllowHeaders + ",x-applet-type"
	response.CORS(option)
	if r.Method == "OPTIONS" {
		r.Response.WriteStatusExit(http.StatusOK)
	} else {
		r.Middleware.Next()
	}
}
