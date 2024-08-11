/*
 * @Author: yanghang
 * @Date: 2021-08-26 10:02:05
 * @LastEditors: yanghang
 * @Description:
 */
package response

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)
var (
	//通用错误码
	ResCodeSuccess = gcode.New(1, "成功", nil)  //成功
	ResCodeError   = gcode.New(-1, "失败", nil) //失败

	//其他通用错误
	ResAuthError         = gcode.New(-100000001, "授权错误", nil) //授权错误
	ResSearchEngineError = gcode.New(-100000002, "搜索引擎错误", nil) //搜索引擎错误
	ResCacheReadError    = gcode.New(-100000003, "缓存获取错误", nil) //缓存获取错误
	ResSmsSendError      = gcode.New(-100000004, "", nil) //
)


// 数据返回通用JSON数据结构
type JsonRes struct {
	Code int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
	Time string      `json:"time"` //当前服务器时间
}

// 返回标准JSON数据。
func Json(r *ghttp.Request, code gcode.Code, msg string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code: code.Code(),
		Msg:  msg,
		Data: responseData,
		Time: gtime.Now().Format("Y-m-d H:i:s"),
	})
}

func JsonError(r *ghttp.Request, err error) {
	switch e := err.(type) {
	case *gerror.Error:
		if e.Code() == ResCodeError {
			g.Log().Errorf(context.Background(),"%+v", err)
		}
		errmsg := e.Error()
		if e.Code().Message() != "" {
			errmsg = e.Code().Message()
		}
		Json(r, e.Code(), errmsg, g.Map{})
	default:
		Json(r, ResCodeError, e.Error(), g.Map{})
	}
	r.Exit()
}

// 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code gcode.Code, msg string, data ...interface{}) {
	Json(r, code, msg, data...)
	r.Exit()
}
