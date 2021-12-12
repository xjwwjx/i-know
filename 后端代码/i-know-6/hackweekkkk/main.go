package main

import (
	"a/database"
	"a/routers"
	"fmt"
)

func main() {
	err := database.Init() //数据库初始化
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	routers.Init() //路由初始化
}
