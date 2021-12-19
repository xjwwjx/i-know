package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
// 连接数据库参数


var DB *sql.DB
// 连接数据库
func initDB() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := "root:xjw2003XJW2021@tcp(localhost:3306)/mdb"
	// 打开数据库 （"驱动名"，连接）
	DB,_ = sql.Open("mysql",path)
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置数据库最大闲置数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	fmt.Println("连接数据库成功")
}

type User struct {
	id int64 //用户名
	password string //密码
}

// @userLogin
func userLogin(c *gin.Context) {
	ID := c.PostForm("id")                // 666666666666666
	passWord := c.PostForm("password")
	//c.HTML(200,"4.html",gin.H{
	//	"Name":userName,
	//	"Password":passWord,
	//})
var s User

	sqlStr:=`SELECT * FROM user where id=?;`
    row:=DB.QueryRow(sqlStr,ID)
   row.Scan(&s.id,&s.password)
	fmt.Printf("调用了%#v\n\n",s)
	if s.id==0 {
			c.JSON(200, gin.H{
				"success": false,
				"code":    400,
				"msg":     "无此用户",
			})
		} else {

				if passWord != s.password {
					c.JSON(200, gin.H{
						"success": false,
						"code":    400,
						"msg":     "密码错误",
					})
				} else {
					c.JSON(200, gin.H{
						"success": true,
						"code":    200,
						"msg":     "登录成功",
					})
				}
			}
		}

func main(){
	initDB()
 	user := gin.Default()
// 	//user := router.Group("")
// 	//{
 		user.POST("/login",userLogin)

// 	//}
	user.Run(":9000")

}
