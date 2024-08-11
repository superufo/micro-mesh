package share

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/micro-mesh/common/response"
)

var ErrorMiddleware = errorMiddleware{}

type errorMiddleware struct {
}

func (s *errorMiddleware) ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	//请求错误统一返回
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		response.Json(r, response.ResCodeError, "系统，请稍后再试吧！")
	}

	//部分错误处理屏蔽显示给前端具体错误
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		response.Json(r, response.ResCodeError, "服务开小差了，请稍后再试吧！！")
		return
	}
}
