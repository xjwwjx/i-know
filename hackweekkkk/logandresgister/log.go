package logandresgister

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"gorm.io/driver/mysql"
)

type source2 struct {
	USERNAME string `form:"username" json:"username" binding:"required"`
	PASSWORD string `form:"password" json:"password" binding:"required"`
}

func Loggg() {
	r := gin.New()
	r.Use(gin.Recovery())
	var message source2
	var check source2
	r.POST("/form", func(c *gin.Context) {
		dsn := "root:xjw2003XJW2021@tcp(127.0.0.1:8081)/xjwwjx?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	r.Run(":8081")
}
