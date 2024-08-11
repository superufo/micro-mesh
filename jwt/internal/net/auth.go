package net

type UserLoginReq struct {
	Account string   `p:"account"  json:"account" v:"required#帐号不能为空"`
	OrgCode string   `p:"org_code" json:"org_code"`
	PassWord string  `p:"password" json:"password" v:"required#密码不能为空"`
	Type     int     `p:"type"     json:"type"`
}

