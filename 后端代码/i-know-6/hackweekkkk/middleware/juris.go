package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdmin() gin.HandlerFunc { //中间件
	return func(ctx *gin.Context) {
		code := ctx.Request.Header.Get("code")
		if code != "iknowyyds" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "无权访问",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
