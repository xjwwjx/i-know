package logandresgister

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type source1 struct {
	USERNAME string `form:"username" json:"username" binding:"required"`
	PASSWORD string `form:"password" json:"password" binding:"required"`
}

func Logandregister() {
	r := gin.New()
	r.Use(gin.Recovery())
	var message source1
	var check source1
	dsn := "root:xjw2003XJW2021@tcp(127.0.0.1:3306)/xjwwjx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	r.POST("/register", func(c *gin.Context) {
		c.ShouldBind(&message)
		if err != nil {
			println("failed to connect database!")
		}
		db.First(&check, "username=?", message.USERNAME)
		if check.USERNAME == "" {
			db.Model(&source1{}).Create(message)
			c.String(http.StatusOK, "success")
		} else {
			c.String(http.StatusOK, fmt.Sprintf("the user already exists !!!"))
		}
	})
	r.POST("/log", func(c *gin.Context) {
		if err != nil {
			println("failed to connect database!")
		}
		message.USERNAME = c.PostForm("username")
		message.PASSWORD = c.PostForm("password")
		db.First(&check, "username=?", message.USERNAME, "password=?", message.PASSWORD)
		if check.USERNAME == "" || check.PASSWORD == "" {
			c.String(http.StatusOK, fmt.Sprintf("the account or password is incorrect !!!"))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("login successfully !!!!!"))
		}
	})
	r.Run("localhost:8080")
}
