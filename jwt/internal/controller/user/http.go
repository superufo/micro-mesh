package user

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/micro-mesh/common/jwt"
	"github.com/micro-mesh/common/response"
	cservice "github.com/micro-mesh/common/service"

	"github.com/micro-mesh/jwt/internal/enum"
	"github.com/micro-mesh/jwt/internal/net"
	"github.com/micro-mesh/jwt/internal/service"
)

type userController struct{}

var User = userController{}

// r *ghttp.Request   ctx context.Context, req *api.AuthLoginReq
func (c *userController) Login(r *ghttp.Request) {
	var (
		in  *net.UserLoginReq
		ctx = r.Context()
	)

	g.Log().Info(ctx, "r:", r)
	g.Log().Info(ctx, "r.Body:", r.Body)
	g.Log().Info(ctx, "in:", in)

	if err := r.Parse(&in); err != nil {
		response.JsonError(r, gerror.WrapCode(enum.RecordNotFindCode, err, ""))
	}

	_, err := service.User.LoginByAccount(ctx, in)
	if err != nil {
		response.JsonError(r, gerror.NewCode(response.ResCodeError, "生成token失败"))
	}

	userContext := cservice.ContextService.Get(r.Context())
	token, err := jwt.GeneUserToken(userContext.User)
	//生成jwt返回

	if err != nil {
		response.JsonError(r, err)
	}

	response.Json(r, response.ResCodeSuccess, "", g.Map{
		"token": token,
	})
}

func (c *userController) Info(r *ghttp.Request) {

}
