package response

import "xxoo/model/po"

type UserVo struct {
	po.SysUser
	RoleId          int64  `json:"roleId"`
	RoleName        string `json:"roleName"`
	RoleDescription string `json:"roleDescription"`
	RoleSeq         int64  `json:"roleSeq"`
	RoleStatus      int64  `json:"roleStatus"`
	RolesList       string `json:"rolesList"`
}
