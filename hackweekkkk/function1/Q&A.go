package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Questions struct {
	ID       int
	QUE      string    `form:"que" json:"que" binding:"required"`
	UserName string    `form:"username" json:"username" binding:"required"`
	Answers  []Answers `gorm:"polymorphic:Question;"`
}

type Answers struct {
	ANS          string `form:"ANS" json:"ANS" binding:"required"`
	UserName     string `form:"UserName" json:"UserName" binding:"required"`
	QuestionID   int
	QuestionType string
}

func main() {
	r := gin.Default()
	dsn := "root:xjw2003XJW2021@tcp(127.0.0.1:3306)/iknow?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("failed to connect database!")
	}
	db.AutoMigrate(&Questions{})
	db.AutoMigrate(&Answers{})
	r.GET("get", func(c *gin.Context) {
	})
}
