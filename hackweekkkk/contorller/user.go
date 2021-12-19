package contorller

import (
	"a/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"

	"a/models"
	"net/http"
)

func Register(c *gin.Context) {
	var message models.UserDatabase
	var check models.UserDatabase
	db := dao.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id=?",message.ID).First(&check)
	if check.ID == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(message.Password), bcrypt.DefaultCost) //加密处理
		err = bcrypt.CompareHashAndPassword([]byte(message.Password), []byte(hash))
		if err != nil {
			fmt.Println(err)
		}
		message.Password = string(hash)
		db.Model(&models.UserDatabase{}).Create(&message)
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "the id already exists",
		})
		check.ID = 0
	}
}
func Login(c *gin.Context) {
	var message models.UserParam
	var check models.UserParam
	//var check1 models.UserDatabase

	db := dao.GetDB()
	err := c.ShouldBind(&message)   //绑定输入的参数
	if err != nil {
		return
	}
	//fmt.Println("?")
	//fmt.Println(check1.Password,check.Password,message.Password)
	//fmt.Println("?")
	//hash,_:= bcrypt.GenerateFromPassword([]byte(message.Password), bcrypt.DefaultCost) //加密处理


//fmt.Println(hash,[]byte(message.Password),err1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//message.Password = string(hash)
	//db.Model(&models.UserDatabase{}).Where("id=? AND password=?", message.ID, message.Password).First(&check)

	db.Model(&models.UserDatabase{}).Where("id=? ", message.ID).First(&check)  //从注册里面的数据库取数据---id password
//fmt.Println(message.Password,hash)
//	fmt.Println(check.Password)
	err1 := bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(message.Password))  //hash与password比较，如果转化后的密码相同则为nil，否则为err

	if check.ID == 0 || err1 != nil {
		c.JSON(200, gin.H{
			"msg": "the id or password is not true",
		})
	} else{
		c.JSON(200, gin.H{
			"msg": "log successfully",
		})
		check.Password = ""
		check.ID = 0
	}
}
func Forget(c *gin.Context) {
	var message models.UserParamForget
	var check models.UserParamForget
	err := c.ShouldBind(&message)
	if err != nil {
		return
	}
	db := dao.GetDB()
	db.Table("UserDatabase").Where("mail=? AND id=?", message.Mail, message.ID).First(&check)
	if check.Mail == "" {
		c.JSON(200, gin.H{
			"msg": "the id or mail is not true",
		})
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(message.Password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			fmt.Println(err)
		}
		check.Password = string(hash)
		db.Save(&check)
		check.Password = ""
		check.ID = 0
		check.Mail = ""
		c.JSON(200, gin.H{
			"msg": "success",
		})
	}
}
func HandQue(c *gin.Context) {
	var message models.QuestionsParam
	db := dao.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Create(&message)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
func HandAns(c *gin.Context) {
	var message models.AnswersParam
	var check models.QuestionDatabase
	db := dao.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		return
	}
	db.Create(&message)
	db.Where("id=?", message.QuestionID).Find(&check)
	check.AnsNum = check.AnsNum + 1
	db.Save(&check)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
func Hot(c *gin.Context) {
	var message []models.QuestionHot
	db := dao.GetDB()
	db.Limit(10).Where("great desc").Find(&message)
	c.JSON(200, gin.H{
		"hot": message,
	})
}
func Search(c *gin.Context) {
	var message string
	var questions []models.QuestionHot
	db := dao.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		return
	}
	db.Where("name like ?", message).Find(&questions)
	c.JSON(200, gin.H{
		"search": questions,
	})
}

func Adore(c *gin.Context) {
	var message models.Like
	var check models.QuestionDatabase
	db := dao.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "ERROR",
		})
		return
	}
	db.Where("id=?", message.Num).Find(&check)
	check.Great = check.Great + 1
	db.Save(&check)
}