package logandresgister

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type users struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	USERNAME string `form:"username" json:"username" binding:"required"`
	PASSWORD string `form:"password" json:"password" binding:"required"`
	MAIL     string `form:"mail" json:"mail" binding:"required"`
}

func LogAndRegister() {
	r := gin.Default()
	var message users
	var check users
	dsn := "root:xjw2003XJW2021@tcp(127.0.0.1:3306)/iknow?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("failed to connect database!")
	}
	db.AutoMigrate(&users{})
	r.POST("/register", func(c *gin.Context) {
		c.ShouldBind(&message)
		db.First(&check, "id=?", message.ID)
		if check.ID == 0 {
			db.Model(&users{}).Create(&message)
			c.String(http.StatusOK, "success")
		} else {
			c.String(http.StatusOK, fmt.Sprintf("the user already exists !!!"))
			check.ID = 0
		}
	})
	r.POST("/log", func(c *gin.Context) {
		if err != nil {
			println("failed to connect database!")
		}
		c.ShouldBind(&message)
		db.Where("id=? AND password=?", message.ID, message.PASSWORD).First(&check)
		if check.ID == 0 || check.PASSWORD == "" {
			c.String(http.StatusOK, fmt.Sprintf("the id or password is incorrect !!!"))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("login successfully !!!!!"))
			check.PASSWORD = ""
			check.ID = 0
		}
	})
	r.PUT("forget", func(c *gin.Context) {
		c.ShouldBind(&message)
		db.Where("mail=?", message.MAIL).First(&check)
		if check.MAIL == "" {
			c.String(http.StatusOK, fmt.Sprintf("the mail is not existed!!!"))
		} else {
			newpassword := message.PASSWORD
			check.PASSWORD = newpassword
			db.Save(&check)
			c.String(http.StatusOK, fmt.Sprintf("success"))
		}
	})
	r.Run(":8080")
}
