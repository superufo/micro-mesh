package user

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/micro-mesh/common/jwt"
	"github.com/micro-mesh/jwt/internal/model"
	"github.com/micro-mesh/jwt/internal/net"
	"github.com/micro-mesh/jwt/internal/service"

	protobuf "github.com/micro-mesh/common/net/pb/user"
	"github.com/micro-mesh/common/response"
	cSrv "github.com/micro-mesh/common/service"
)

var gctx context.Context

// var UserGrpcController = userGrpcController{}  Token(context.Context, *UserLoginTos) (*ReplyToc, error)  pb.UnimplementedUserServer pb.UnimplementedUserServer
func Register(s *grpcx.GrpcServer, ctx context.Context) {
	gctx = ctx

	protobuf.RegisterUserServer(s.Server, &GrpcController{})
}

type GrpcController struct {
	protobuf.UnimplementedUserServer
}

func (s GrpcController) mustEmbedUnimplementedUserServer() {

}

// GetToken
func (s GrpcController) GetToken(ctx context.Context, input *protobuf.UserLoginTos) (*protobuf.ReplyToc, error) {
	ctx = gctx
	g.Log().Debugf(ctx, "input:", input)

	var (
		in  = &net.UserLoginReq{}
		res = &protobuf.ReplyToc{}
		err = errors.New("")
	)

	{
		in.Account = string(input.Account)
		in.PassWord = input.Password
		if input.OrgCode != nil {
			in.OrgCode = *input.OrgCode
		}
		if input.Type != nil {
			in.Type = int(*input.Type)
		}
	}

	// 获取用户信息
	userContext, err := service.User.LoginByAccount(ctx, in)
	//g.Log().Info(ctx,"userContext:", cservice.ContextService.Get(ctx))
	if err != nil {
		res.Code = uint32(response.ResAuthError.Code())
		res.Msg = "认证失败了" //response.ResAuthError.Message()
		res.Data = nil

		return res, err
	}

	//生成token
	//userContext := cservice.ContextService.Get(ctx)
	token, err := jwt.GeneUserToken(userContext)
	if err != nil {
		res.Code = uint32(response.ResAuthError.Code())
		res.Msg = err.Error()
		res.Data = nil
		return res, err
	}

	res.Code = uint32(response.ResCodeSuccess.Code())
	res.Msg = ""
	res.Data = &protobuf.DataToc{
		Token: token,
		Time:  gtime.Now().Format("Y-m-d H:i:s"),
	}

	return res, nil
}

func (s GrpcController) VerifyToken(ctx context.Context, input *protobuf.VerifyTokenTos) (*protobuf.VerifyTokenToc, error) {
	var (
		res = &protobuf.VerifyTokenToc{}
		err = errors.New("")

		context = cSrv.ContextService
	)

	authorization := input.Token
	ContextUser, err := jwt.ParseUserToken(authorization)
	if err == nil {
		//判断用户的登录是否过期
		nowTime := time.Now().Unix()
		if nowTime > ContextUser.ExpiresAt {
			res.Code = uint32(response.ResAuthError.Code())
			res.Msg = "token已经过期"
			res.Data = false

			return res, err
		}

		//开始判断用户和机构状态是否正常,全部正常的情况写入上下文
		var user *model.SysUser
		err = g.Model("sys_user").Where("user_id=?", ContextUser.UserId).Scan(&user)
		if err == nil && user != nil && user.Status == 1 {
			if user.UpdatedAt.Time.Unix() > ContextUser.CreatedAt {
				res.Code = uint32(response.ResAuthError.Code())
				res.Msg = "用户信息已更改，请重新登录"
				res.Data = false

				return res, err
			}
			//查找归属机构
			var userOrg *model.BaseOrg
			err = g.Model("base_org").Where("org_id=?", user.OrgId).Scan(&userOrg)
			if err == nil && userOrg != nil && userOrg.Status == 2 {
				context.SetUser(ctx, ContextUser)
			}
		}
	} else {
		res.Code = uint32(response.ResAuthError.Code())
		res.Msg = "登录过期，请重新登录"
		res.Data = false

		return res, err
	}

	res.Code = uint32(response.ResAuthError.Code())
	res.Msg = "token验证通过"
	res.Data = true

	return res, err
}
