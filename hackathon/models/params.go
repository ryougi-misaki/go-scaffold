package models

type ParamRegister struct {
	Name      string `json:"name" form:"name"`
	Telephone string `json:"telephone" form:"telephone" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Telephone string `json:"telephone" form:"telephone" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}
