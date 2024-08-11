package model

import "github.com/gogf/gf/v2/os/gtime"

type SysUserAccount struct {
	UserId    uint        `orm:"user_id"    json:"userId"    description:""`                   //
	Type      uint        `orm:"type"       json:"type"      description:"1 用户名，2手机号码， 3 会员号"` // 1 用户名，2手机号码， 3 会员号
	Account   string      `orm:"account"    json:"account"   description:"账号"`                 // 账号
	Status    uint        `orm:"status"     json:"status"    description:"状态"`                 // 状态
	LeaveDate *gtime.Time `orm:"leave_date" json:"leaveDate" description:"离职时间"`               // 离职时间
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt" description:"创建时间"`               // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt" description:"更新时间"`               // 更新时间
}

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	UserId      uint        `orm:"user_id,primary" json:"userId"      description:""`                                   //
	Mobile      string      `orm:"mobile"          json:"mobile"      description:"手机号码"`                               // 手机号码
	UserName    string      `orm:"user_name"       json:"userName"    description:"帐号"`                                 // 帐号
	Name        string      `orm:"name"            json:"name"        description:"姓名"`                                 // 姓名
	Status      uint        `orm:"status"          json:"status"      description:"状态 0 禁用 1 启用"`                       // 状态 0 禁用 1 启用
	OrgId       uint        `orm:"org_id"          json:"orgId"       description:"机构id,公司归属"`                          // 机构id,公司归属
	IsSystem    uint        `orm:"is_system"       json:"isSystem"    description:"是否系统管理员,系统管理员禁止修改和拥有所有权限  0 否  1 是"` // 是否系统管理员,系统管理员禁止修改和拥有所有权限  0 否  1 是
	IsPredefine uint        `orm:"is_predefine"    json:"isPredefine" description:"是否预置帐号(0 否 1 是),预置帐号不能修改信息"`         // 是否预置帐号(0 否 1 是),预置帐号不能修改信息
	CreatedAt   *gtime.Time `orm:"created_at"      json:"createdAt"   description:"创建时间"`                               // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"      json:"updatedAt"   description:"更新时间"`                               // 更新时间
	DeletedAt   *gtime.Time `orm:"deleted_at"      json:"deletedAt"   description:"删除时间"`                               // 删除时间
}

// SysUserPassword is the golang structure for table sys_user_password.
type SysUserPassword struct {
	UserId    uint        `orm:"user_id,primary" json:"userId"    description:""`       //
	Type      uint        `orm:"type,primary"    json:"type"      description:"1 登录密码"` // 1 登录密码
	Password  string      `orm:"password"        json:"password"  description:"密码"`     // 密码
	CreatedAt *gtime.Time `orm:"created_at"      json:"createdAt" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at"      json:"updatedAt" description:"更新时间"`   // 更新时间
}

// BaseOrg is the golang structure for table base_org.
type BaseOrg struct {
	OrgId         uint        `orm:"org_id,primary"  json:"orgId"         description:"机构id"`                             // 机构id
	Pid           uint        `orm:"pid"             json:"pid"           description:"上级机构"`                             // 上级机构
	OrgCode       string      `orm:"org_code"        json:"orgCode"       description:"机构代码"`                             // 机构代码
	OrgName       string      `orm:"org_name"        json:"orgName"       description:"机构名称"`                             // 机构名称
	OrgTypeId     uint        `orm:"org_type_id"     json:"orgTypeId"     description:"机构类型  1：公司(平台)，2：区域，3：代理商，4：家政企业"` // 机构类型  1：公司(平台)，2：区域，3：代理商，4：家政企业
	HasChildren   uint        `orm:"has_children"    json:"hasChildren"   description:"是否存在子节点"`                          // 是否存在子节点
	ServiceAreaId string      `orm:"service_area_id" json:"serviceAreaId" description:""`                                 //
	Status        int         `orm:"status"          json:"status"        description:"状态（0：草稿，1：待审核，2：生效，9：失效）"`         // 状态（0：草稿，1：待审核，2：生效，9：失效）
	Desc          string      `orm:"desc"            json:"desc"          description:"描述"`                               // 描述
	CreatedAt     *gtime.Time `orm:"created_at"      json:"createdAt"     description:"创建时间"`                             // 创建时间
	UpdatedAt     *gtime.Time `orm:"updated_at"      json:"updatedAt"     description:"更新时间"`                             // 更新时间
	DeletedAt     *gtime.Time `orm:"deleted_at"      json:"deletedAt"     description:"删除时间"`                             // 删除时间
}








