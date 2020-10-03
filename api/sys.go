package api

import (
	list2 "container/list"
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	coven2 "github.com/petersunbag/coven"
	"net/http"
	"time"
	"xxoo/config"
	"xxoo/model"
	"xxoo/model/po"
	"xxoo/model/request"
	"xxoo/model/response"
	"xxoo/service"
	"xxoo/utils"
)

func GetUserList(c *gin.Context) {
	var userQueryVO request.UserQueryVO
	c.BindJSON(&userQueryVO)

	var page = new(model.PageModel)
	page.List, page.TotalCount = service.GetUserList(userQueryVO)
	page.CurrPage = userQueryVO.Page
	page.PageSize = userQueryVO.PageSize
	resultMap := map[string]interface{}{}
	resultMap["page"] = page
	response.OkWithData(resultMap, c)
}

func DeleteUserById(c *gin.Context) {
	var idRequest request.IdRequest
	c.BindJSON(&idRequest)
	service.DeleteUser(idRequest.ID)
	c.JSON(http.StatusOK, model.ResultMap{Msg: "success"})
}

func Update(c *gin.Context) {
	var user po.SysUser
	c.BindJSON(&user)
	user.UpdateTime = time.Now()
	user.CreateTime = time.Now()
	service.Update(user)
	c.JSON(http.StatusOK, model.ResultMap{Msg: "success"})
}

func AddUser(c *gin.Context) {
	var user po.SysUser
	c.BindJSON(&user)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	service.AddUser(user)

	c.JSON(http.StatusOK, model.ResultMap{Msg: "success"})
}

func Login(c *gin.Context) {
	var L request.RegisterAndLoginStruct
	c.ShouldBindJSON(&L)
	RuleMap := utils.Rules{"UserName": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}}
	if VerifyError := utils.Verify(L, RuleMap); VerifyError != nil {
		response.FailWithMessage(VerifyError.Error(), c)
		return
	}

	U := &po.SysUser{UserName: L.Username, Password: L.Password}
	err, user := service.Login(U)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	generateToken(c, *user)

}

func generateToken(c *gin.Context, user po.SysUser) {
	StandardClaims := jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
		ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
		Issuer:    "xoplus",                       // 签名的发行者
	}

	token, err := createToken(StandardClaims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userJson, _ := json.Marshal(user)
	config.REDIS.Set(token, userJson, 1000*1000*1000*60*60*24)
	response.OkWithData(map[string]interface{}{
		"user":      user,
		"token":     token,
		"expiresAt": StandardClaims.ExpiresAt * 1000,
	}, c)

}

func createToken(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SIGNING_KEY))
}

func Nav(c *gin.Context) {
	token := c.Request.Header.Get("token")

	userinfoStr, _ := config.REDIS.Get(token).Result()
	userinfo := po.SysUser{}
	json.Unmarshal([]byte(userinfoStr), &userinfo)

	l := service.GetUserRoles(userinfo.ID)
	subMenuList := list2.New()
	resultMap := map[string]interface{}{}
	menuSet := mapset.NewSet()
	permissions := mapset.NewSet()
	var sysResourceList []po.SysResource
	for i := 0; i < len(l); i++ {
		sysResourceList = service.GetRoleResources(l[i])

		for i := 0; i < len(sysResourceList); i++ {
			value := (sysResourceList)[i]
			if value.ParentId == 0 {
				subMenuList.PushFront(value)
			}
		}
	}

	for value := subMenuList.Front(); value != nil; value = value.Next() {
		parentResource := value.Value.(po.SysResource)
		recursionMenu(permissions, menuSet, &parentResource, sysResourceList)
	}
	resultMap["menuList"] = menuSet
	resultMap["permissions"] = permissions
	resultMap["userinfo"] = userinfo
	response.OkWithData(resultMap, c)

}

func recursionMenu(permissions mapset.Set, menuSet mapset.Set, parentResource *po.SysResource, sysResourceList []po.SysResource) {
	subResourceList := []po.SysResource{}
	for i := 0; i < len(sysResourceList); i++ {

		sr := (sysResourceList)[i]
		if sr.ParentId == parentResource.ID {
			subResourceList = append(subResourceList, sr)
		}
	}

	if len(subResourceList) > 0 {
		subMenuSlice := []po.SysResource{}
		if parentResource.SubMenuList == nil {
			parentResource.SubMenuList = subMenuSlice
		}
		for i := 0; i < len(parentResource.SubMenuList); i++ {
			subMenuSlice = append(subMenuSlice, (parentResource.SubMenuList)[i])
		}
		subMenuSlice = append(subMenuSlice, subResourceList...)
		if permissions != nil {
			for i := 0; i < len(subResourceList); i++ {
				permissions.Add(subResourceList[i].Url)
			}
		}

		parentResource.SubMenuList = subMenuSlice

		for i := 0; i < len(parentResource.SubMenuList); i++ {
			recursionMenu(permissions, menuSet, &parentResource.SubMenuList[i], sysResourceList)
		}

		if parentResource.ParentId == 0 {
			menuSet.Add(parentResource)
		}

	}

}

func GetRoles(c *gin.Context) {
	var page = new(model.PageModel)
	page.List = service.GetRoles()
	resultMap := map[string]interface{}{}
	resultMap["page"] = page
	response.OkWithData(resultMap, c)
}

func RoleInfo(c *gin.Context) {
	var idRequest request.IdRequest
	c.ShouldBindJSON(&idRequest)

	subMenuList := list2.New()
	resultMap := map[string]interface{}{}
	menuSet := mapset.NewSet()
	var allResourceList []po.SysResource
	allResourceList = service.GetAllResources()

	for i := 0; i < len(allResourceList); i++ {
		value := (allResourceList)[i]
		if value.ParentId == 0 {
			subMenuList.PushFront(value)
		}
	}

	for value := subMenuList.Front(); value != nil; value = value.Next() {
		parentResource := value.Value.(po.SysResource)
		recursionMenu(nil, menuSet, &parentResource, allResourceList)
	}

	roleResources := service.GetRoleResources(idRequest.ID)

	for elem := range menuSet.Iterator().C {
		currentResource := elem.(*po.SysResource)
		currentResource.CheckStatus = 2
		if currentResource.SubMenuList != nil {
			recursionCheckMenuStatus(currentResource.SubMenuList, roleResources)
		} else {
			for i := 0; i < len(roleResources); i++ {
				if currentResource.ID == roleResources[i].ID {
					currentResource.CheckStatus = 1
					break
				}
			}
		}

	}

	resultMap["role"] = menuSet
	response.OkWithData(resultMap, c)

}
func recursionCheckMenuStatus(subMenuList []po.SysResource, roleResource []po.SysResource) {

	for i := 0; i < len(subMenuList); i++ {
		subMenuList[i].CheckStatus = 2
		if subMenuList[i].SubMenuList != nil {
			recursionCheckMenuStatus(subMenuList[i].SubMenuList, roleResource)
		} else {
			for j := 0; j < len(roleResource); j++ {
				if subMenuList[i].ID == roleResource[j].ID {
					subMenuList[i].CheckStatus = 1
					break
				}
			}
		}
	}

}

func DeleteRoleById(c *gin.Context) {
	var idRequest request.IdRequest
	c.BindJSON(&idRequest)
	service.DeleteRoleResource(idRequest.IDS)
	service.DeleteRoleById(idRequest.IDS)
	response.Ok(c)
}

func UpdateRole(c *gin.Context) {
	var roleVo request.RoleVo
	c.BindJSON(&roleVo)
	if roleVo.ID != 0 {
		service.DeleteRoleResource([]int64{roleVo.ID})
		var c, _ = coven2.NewConverter(po.SysRole{}, request.RoleVo{})
		role := po.SysRole{}
		c.Convert(&role, &roleVo)
		service.UpdateRole(role)

	}
	resourceList := roleVo.ResourceIdList
	roleResourceList := []po.SysRoleResource{}
	for i := 0; i < len(resourceList); i++ {
		resourceId := resourceList[i]
		roleResourceList = append(roleResourceList, po.SysRoleResource{RoleId: roleVo.ID, ResourceId: resourceId})
	}

	service.AddRoleResources(roleResourceList)
	response.Ok(c)
}
