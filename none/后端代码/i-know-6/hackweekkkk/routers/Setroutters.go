package routers

import (
	"a/contorller"
	"a/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r = CollectRoute(r)
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("fail to start gin,err:%v\n", err)
		return
	}
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("register", contorller.Register)
	r.POST("log", contorller.Login)
	r.PUT("forget", contorller.Forget)
	r.POST("hand_que", contorller.HandQue)
	r.POST("hand_ans", middleware.AuthAdmin(), contorller.HandAns)
	r.GET("hot", contorller.Hot)
	r.POST("search", contorller.Search)
	r.POST("like", contorller.Adore)
	r.POST("collect", contorller.Collect)
	return r
}
