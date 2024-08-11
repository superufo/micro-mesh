package model

import "github.com/gogf/gf/v2/frame/g"

const (
	ContextKey = "AppData"
)

// 会员请求上下文结构
type MemberContext struct {
	Member *ContextMember // 上下文用户信息
	Config *ContextConfig //上下文配置
	Data   g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// 请求上下文中的用户信息
type ContextMember struct {
	UserId    uint  `json:"userId"`    // 用户ID
	Status    int   `json:"status"`    // 用户状态
	UserType  uint  `json:"userType"`  //用户类型
	ExpiresAt int64 `json:"expiresAt"` //失效时间
}

// 请求上下文中的用户信息
type ContextConfig struct {
	AppletType string // 小程序类型
}

//后台用户请求上下文
type UserContext struct {
	User *ContextUser // 上下文用户信息
	Data g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}

// 请求上下文中的用户信息
type ContextUser struct {
	UserId    uint  `json:"userId"`    // 用户ID
	OrgId     uint  `json:"orgId"`     // 机构id 公司级
	CreatedAt int64 `json:"createdAt"` //token生成时间
	ExpiresAt int64 `json:"expiresAt"` //失效时间
}

