package request


type RoleVo struct {
	ID             int64   `json:"id"`
	Description    string  `json:"description"`
	Name           string  `json:"name"`
	Seq            int     `json:"seq"`
	ResourceIdList []int64 `json:"resourceIdList"`
}

type RegisterAndLoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserQueryVO struct {
	PageRequest
	UserName string `json:"userName"`
	RoleId   int64  `json:"roleId"`
}


