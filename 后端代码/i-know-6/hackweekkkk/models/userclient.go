package models

type UserParamLog struct {
	ID       uint   `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type UserParamForget struct {
	ID       uint   `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Mail     string `form:"mail" json:"mail" binding:"required"`
}
type UserParamRegister struct {
	ID       uint   `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Mail     string `form:"mail" json:"mail" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
}
