package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Expose-Headers", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}


func AuthAdmin() gin.HandlerFunc { //中间件
	return func(ctx *gin.Context) {
		code := ctx.Request.Header.Get("code")
		if code != "iknow" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "无权访问（密码是产品名字）",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
