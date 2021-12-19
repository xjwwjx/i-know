package main

import (
	"a/dao"
	"a/routers"
)

func main() {
	//err := dao.Init()
	//if err != nil {
	//	fmt.Printf("init mysql failed, err:%v\n", err)
	//	return
	//}
	dao.Init()
	routers.Init()
}
