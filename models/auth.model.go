package models

type Login struct {
	UserName string `json:"userName" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
