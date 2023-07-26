package dto

type LoginForm struct {
	Id       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type LoginJSON struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
