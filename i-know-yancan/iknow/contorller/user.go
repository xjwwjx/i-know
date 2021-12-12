package contorller

import (
	"a/database"
	util "a/jwt"
	"strings"

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
		//err = bcrypt.CompareHashAndPassword([]byte(message.Password), []byte(hash))
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
			"msg": "register successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "the id already exists",
		})
		check.ID = 0
		check.Username = ""
	}
}
func Login(c *gin.Context) {
	var message models.UserParamLogin
	var check models.UserParamLogin
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
			"msg": "login successfully",
		})
		check.Password = ""
		check.ID = 0
//------------------------------------------------------
		token,err:=util.GeterateToken(message.ID)
		if err!=nil{
			fmt.Println(err)
			c.JSON(200,gin.H{
				"msg":"创建token失败",
			})
		}
		c.JSON(200,gin.H{
			"msg":"创建token成功",
			"Token":token,
		})
		return
	}
	c.JSON(200,gin.H{
		"msg":"鉴权失败",
	})


//-----------------------------------------------
}
func DelId(c *gin.Context){ //删除已经注册的id
	var message models.UserParamDelete
	//var delete models.QuestionDatabase
	db :=database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&models.UserDatabase{}).Where("id=? ", message.ID).Delete(&message.ID)
	c.JSON(200, gin.H{
		"msg": "delete successfully",
	})
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
		hash, err := bcrypt.GenerateFromPassword([]byte(message.NewPassword), bcrypt.DefaultCost) //加密处理
		if err != nil {
			fmt.Println(err)
		}
		check.Password = string(hash)
		db.Save(&check)
		check.Password = "" //清缓存
		check.ID = 0
		check.Mail = ""
		c.JSON(200, gin.H{
			"msg": "reset password successfully",
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
		"msg": "handle question successfully",
	})
}
func DelQue(c *gin.Context){ //根据id删除提交的问题
	var message models.DelQuestionsParam
	//var delete models.QuestionDatabase
	db :=database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&models.QuestionDatabase{}).Where("id=? ", message.ID).Delete(&message.ID)
	c.JSON(200, gin.H{
		"msg": "delete question successfully",
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
		"msg": "hand answer successfully",
	})
}
func DelAns(c *gin.Context){//根据id删除提交的问题所对应的答案
	var message models.DelAnswersParam
	//var delete models.QuestionDatabase
	db :=database.GetDB()
	err := c.ShouldBind(&message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&models.AnswerDatabase{}).Where("id=? ", message.ID).Delete(&message.ID)
	c.JSON(200, gin.H{
		"msg": " delete answer successfully",
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
	c.JSON(200,gin.H{
		"msg":"点赞成功",
	})
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
	c.JSON(200, gin.H{
		"msg": "collect successfully",
	})

}
//-------------------------------------   /JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("id", mc.ID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
func HomeHandler(c *gin.Context) {
	id := c.MustGet("id").(uint)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "verify successfully",
		"id":id,
	})
}
