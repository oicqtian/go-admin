package service

import (
	"errors"
	"xxoo/config"
	"xxoo/model/po"
	"xxoo/model/request"
	"xxoo/model/response"
)

func GetUserList(userQueryVO request.UserQueryVO) (userVoList []response.UserVo, total int64) {
	var userList []response.UserVo

	db := config.DB.Raw("SELECT DISTINCT  t.id,   t.user_name ,     t.`name`,    t.`password`,   t.phone,      t.email,        t.`status`,       t.create_time AS createTime,        t.update_time AS updateTime,t.remark,  t.create_user_name,t3.id AS roleId,t3.`name` AS roleName, t3.description AS roleDescription,t3.seq AS roleSeq,t3.`status` AS roleStatus,group_concat( t3.`id`) AS rolesList  FROM t_sys_user t LEFT JOIN t_sys_user_role t2 ON t.id = t2.user_id   LEFT JOIN t_sys_role t3 ON t2.role_id = t3.id ")
	if userQueryVO.UserName != "" {
		db = db.Where("user_name like ?", "%"+userQueryVO.UserName+"%")
	}
	if userQueryVO.RoleId > 0 {
		db = db.Where("role_id = ? ", userQueryVO.RoleId)
	}
	db.Find(&userList)
	var totalCount int64
	return userList, totalCount
}

func DeleteUser(id int64) (err error) {
	var user po.SysUser
	err = config.DB.Where("id = ?", id).Delete(&user).Error
	return err
}

func Update(user po.SysUser) error {
	e := config.DB.Save(&user).Error
	return e
}

func AddUser(user po.SysUser) error {
	e := config.DB.Create(&user).Error
	return e
}

func Login(u *po.SysUser) (err error, us *po.SysUser) {
	var user po.SysUser
	err = config.DB.Where("user_name=?", u.UserName).First(&user).Error
	if u.Password != user.Password {
		return errors.New("账号密码错误"), nil
	}
	return err, &user
}

func GetUserRoles(userId int64) []int64 {
	var roleIdList []int64
	config.DB.Table("t_sys_user_role").Select("role_id").Where("user_id = ?", userId).Pluck("role_id", &roleIdList)
	return roleIdList
}

func GetRoleResources(roleId int64) []po.SysResource {
	var resourceList []po.SysResource
	config.DB.Raw("select r.* from t_sys_resource r,t_sys_role_resource rr where rr.role_id =? and r.id = rr.resource_id order by r.seq asc", roleId).Find(&resourceList)
	return resourceList
}

func GetAllResources() []po.SysResource {
	var resourceList []po.SysResource
	config.DB.Raw("select * from t_sys_resource order by seq asc  ").Find(&resourceList)
	return resourceList
}

func GetResourceByParentId(ids []int64) []po.SysResource {
	var resourceList []po.SysResource
	config.DB.Table(" t_sys_resource").Where("id in (?)", ids).Find(&resourceList)
	return resourceList
}

func GetRoles() (list interface{}) {
	var roleList []po.SysRole
	config.DB.Find(&roleList)
	return roleList
}

func DeleteRoleById(ids []int64) {
	config.DB.Delete(po.SysRole{}, "id in (?)", ids)
}

func DeleteRoleResource(ids []int64) {
	config.DB.Delete(po.SysRoleResource{}, " role_id  in (?) ", ids)
}

func UpdateRole(role po.SysRole) {
	config.DB.Exec("update t_sys_role set name = ?,seq= ? ,description= ? where id = ?", role.Name, role.Seq, role.Description, role.ID)
}

func AddRoleResources(sysRoleResources []po.SysRoleResource) {
	config.DB.Model(&po.SysRoleResource{}).Select("RoleId", "ResourceId").Create(sysRoleResources)

}
