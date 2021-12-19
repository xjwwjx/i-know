package contorller

import (
	"a/database"

	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"

	"a/models"
	"net/http"
)

func Register(c *gin.Context) {
	var message models.UserParamRegister
	var check models.UserDatabase
	db := database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id=? AND username=?", message.ID, message.Username).First(&check)
	if check.ID == 0 && check.Username == "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(message.Password), bcrypt.DefaultCost) //加密处理
		err = bcrypt.CompareHashAndPassword([]byte(message.Password), []byte(hash))
		if err != nil {
			fmt.Println(err)
		}
		message.Password = string(hash)
		check.ID = message.ID
		check.Username = message.Username
		check.Mail = message.Mail
		check.Password = message.Password
		db.Create(&check) //存新账号进去
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "the id or username already exists",
		})
		check.ID = 0
		check.Username = ""
	}
}
func Login(c *gin.Context) {
	var message models.UserParamLog
	var check models.UserParamLog
	db := database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		return
	}
	db.Model(&models.UserDatabase{}).Where("id=? ", message.ID).First(&check)               //从注册里面的数据库取数据---id password
	err1 := bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(message.Password)) //hash与password比较，如果转化后的密码相同则为nil，否则为err

	if check.ID == 0 || err1 != nil {
		c.JSON(200, gin.H{
			"msg": "the id or password is not true",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "log successfully",
		})
		check.Password = ""
		check.ID = 0
	}
}
func Forget(c *gin.Context) {
	var message models.UserParamForget
	var check models.UserDatabase
	err := c.ShouldBind(&message)
	if err != nil {
		return
	}
	db := database.GetDB()
	db.Where("mail=? AND id=?", message.Mail, message.ID).First(&check)
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
		check.Password = "" //清缓存
		check.ID = 0
		check.Mail = ""
		c.JSON(200, gin.H{
			"msg": "success",
		})
	}
}
func HandQue(c *gin.Context) { //提交问题
	var message models.QuestionsParam
	var upload models.QuestionDatabase
	db := database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	upload.QUE = message.QUE
	upload.UserName = message.UserName
	db.Model(&models.QuestionDatabase{}).Create(&upload)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
func HandAns(c *gin.Context) { //提交答案，需要中间件（权限）
	var message models.AnswersParam
	var upload models.AnswerDatabase
	var check models.QuestionDatabase
	db := database.GetDB()
	err := c.ShouldBind(&message)
	upload.ANS = message.ANS
	upload.UserName = message.UserName
	upload.QuestionID = message.QuestionID
	if err != nil {
		return
	}
	db.Create(&upload)
	db.Where("id=?", message.QuestionID).Find(&check)
	check.AnsNum = check.AnsNum + 1
	db.Save(&check)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
func Hot(c *gin.Context) { //热搜，根据点赞数量
	var message []models.QuestionDatabase
	var show [10]models.QuestionHot
	db := database.GetDB()
	db.Limit(10).Order("great desc").Find(&message)
	for i := 0; i < len(message); i++ {
		show[i].UserName = message[i].UserName
		show[i].Great = message[i].Great
		show[i].QUE = message[i].QUE
		show[i].AnsNum = message[i].AnsNum
	}
	c.JSON(200, gin.H{
		"msg": show,
	})
}
func Search(c *gin.Context) { //从数据库里关键词搜索
	var questions []models.QuestionDatabase
	var upload [10]models.QuestionHot
	var message models.Search
	db := database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "error",
		})
		return
	}
	db.Where("que LIKE ?", "%"+message.KeyWords+"%").Find(&questions)
	for i := 0; i < len(questions); i++ {
		if questions[i].QUE == "" {
			break
		}
		upload[i].QUE = questions[i].QUE
		upload[i].Great = questions[i].Great
		upload[i].AnsNum = questions[i].AnsNum
		upload[i].UserName = questions[i].UserName
	}
	c.JSON(200, gin.H{
		"search": upload,
	})
}

func Adore(c *gin.Context) { //点赞功能
	var message models.Like
	var check models.QuestionDatabase
	db := database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "error",
		})
	}
	db.Where("id=?", message.QuestionID).Find(&check)
	check.Great = check.Great + 1
	db.Model(&check).Update("great", check.Great)
}

func Collect(c *gin.Context) { //收藏功能
	var message models.CollectParam
	var check models.CollectDatabase
	db := database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "error",
		})
	}
	check.UserID = message.UserID
	check.QuestionID = message.QuestionID
	db.Create(&check)
}
