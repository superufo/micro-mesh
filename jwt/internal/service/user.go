package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
	cModel "zhugedaojia.com/common/model"
	"zhugedaojia.com/common/response"
	"zhugedaojia.com/common/utils"

	"zhugedaojia.com/jwt/internal/model"
	"zhugedaojia.com/jwt/internal/net"
)

var User = userService{}

type userService struct{}

func (s *userService) LoginByAccount(ctx context.Context, in *net.UserLoginReq) (*cModel.ContextUser, error) {
	var (
		userAccount *model.SysUserAccount
		userPassword *model.SysUserPassword
		user *model.SysUser
		userOrg *model.BaseOrg
		err error

		password = in.PassWord
		realAccount = in.Account
		orgCode = in.OrgCode
	)
	if orgCode != "" {
		realAccount = orgCode + "_" + realAccount
	}

	{
		err = g.Model("sys_user_account").Where("account=? and leave_date is null", realAccount).Scan(&userAccount)
		if err != nil {
			return nil, gerror.NewCode(response.ResAuthError,"账号查询错误")
		}
		if userAccount == nil {
			return nil,gerror.NewCode(response.ResAuthError, "账号不存在")
		}
		if userAccount.Status == 0 {
			return nil,gerror.NewCode(response.ResAuthError, "账号已停用")
		}
	}

	//判断用户是否存在
	{
		err = g.Model("sys_user").Where("user_id=?", userAccount.UserId).Scan(&user)
		if err != nil {
			return nil,gerror.NewCode(response.ResAuthError, "用户不存在")
		}
		if user.Status == 0 {
			return nil,gerror.NewCode(response.ResAuthError, "用户已停用")
		}
	}

	//检验密码
	{
		err = g.Model("sys_user_password").Where("user_id=? and type = 1", user.UserId).Scan(&userPassword)
		if err != nil {
			return nil,gerror.NewCode(response.ResAuthError, "未设置密码")
		}
		if ok := utils.PasswordVerify(password, userPassword.Password); !ok {
			return nil,gerror.NewCode(response.ResAuthError, "密码错误")
		}
	}

	//查找归属机构
	{
		err = g.Model("base_org").Where("org_id=?", user.OrgId).Scan(&userOrg)
		if err != nil {
			return nil,gerror.NewCode(response.ResAuthError, "机构不存在")
		}

		if userOrg.Status != 2 {
			return nil,gerror.NewCode(response.ResAuthError, "机构已停用")
		}
	}

	contextUser := &cModel.ContextUser{
		UserId:    user.UserId,
		OrgId:     user.OrgId,
		CreatedAt: gtime.Now().Timestamp(),
		ExpiresAt: gtime.Now().Add(12 * time.Hour).Timestamp(),
	}
	g.Log().Info(ctx,"contextUser:", contextUser)

	//设置用户上下文
	//service.ContextService.SetUser(ctx, contextUser)
	return 	contextUser, err
}