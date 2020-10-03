package po

import "xxoo/model"

type SysUser struct {
	model.BaseModel
	UserName       string `json:"userName"`
	Name           string `json:"name"`
	Password       string `json:"-"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Status         int    `json:"status"`
	CreateUserId   int64  `json:"createUserId"`
	Remark         string `json:"remark"`
	CreateUserName string `json:"createUserName"`
	ParentId       int64  `json:"parentId"`
}

type SysRole struct {
	model.BaseModel
	Name         string `json:"name"`
	Seq          int8   `json:"seq"`
	Description  string `json:"description"`
	Status       int8   `json:"status"`
	CreateUserId int64  `json:"createUserId"`
	UserType     int16  `json:"userType"`
}

type SysResource struct {
	model.BaseModel
	ParentId     int64         `json:"parentId"`
	Name         string        `json:"name"`
	Url          string        `json:"url"`
	Description  string        `json:"description"`
	Icon         string        `json:"icon"`
	Seq          int8          `json:"seq"`
	ResourceType int8          `json:"resourceType"`
	Status       int8          `json:"status"`
	SubMenuList  []SysResource `json:"list" gorm:"-"`
	CheckStatus  int           `json:"checkStatus"`
}

type SysRoleResource struct {
	RoleId     int64 `json:"roleId"`
	ResourceId int64 `json:"resourceId"`
}
