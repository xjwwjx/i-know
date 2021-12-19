package routers

import (
	"a/contorller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":8080")
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	//用户
	r.POST("register", contorller.Register)
	r.POST("log", contorller.Login)
	r.PUT("forget", contorller.Forget)
	r.POST("hand_que", contorller.HandQue)
	r.POST("hand_ans", contorller.HandAns)
	r.GET("hot", contorller.Hot)
	r.GET("search", contorller.Search)
	r.POST("like",contorller.Adore)
	return r
}
