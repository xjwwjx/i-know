package models

import "gorm.io/gorm"

type UserDatabase struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;"`
	Password string `gorm:"varchar(25);not null;"`
	Mail     string `gorm:"varchar(25);not null;"`
}
type UserParam struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type UserParamForget struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Mail     string `form:"mail" json:"mail" binding:"required"`
}
type UserParamLog struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Mail     string `form:"mail" json:"mail" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
}
