package router

import (
	"github.com/gin-gonic/gin"
	"xxoo/api"
)

func InitSysRouter(Router *gin.RouterGroup) {
	SysRouter := Router.Group("sys")
	{
		SysRouter.POST("user/list", api.GetUserList)
		SysRouter.POST("deleteById", api.DeleteUserById)
		SysRouter.POST("addUser", api.AddUser)
		SysRouter.POST("update", api.Update)
		SysRouter.POST("login", api.Login)

		SysRouter.POST("menu/nav", api.Nav)

		SysRouter.POST("role/list", api.GetRoles)
		SysRouter.POST("role/info", api.RoleInfo)
		SysRouter.POST("role/update", api.UpdateRole)
		SysRouter.POST("role/deleterolebyid", api.DeleteRoleById)

	}
}
