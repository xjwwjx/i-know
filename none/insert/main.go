package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init()
)
var db *sql.DB
func initDB()(err error){
	//数据库信息
	//"用户名:密码@tcp(ip:端口)/数据库名字"
	dsn := "root:xjw2003XJW2021@tcp(localhost:3306)/mdb"
	//连接数据库
	db,err=sql.Open("mysql",dsn)//不会检验用户名和密码是否正确
	if err != nil{//dsn格式不正确的时候会报错
		return
	}
	err=db.Ping()  //尝试连接数据库    可以检验用户名和密码是否正确
	if err !=nil{
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(100000)
	return
}
type User struct {

	password string

}


func insert(password string){
	err:=initDB()
	if err != nil{
		fmt.Printf("init DB failed,err:%v\n",err)
	}
	fmt.Println("数据库连接成功！")

	sqlStr:=`insert into user (password) values (?);`
	ret,err:=db.Exec(sqlStr,password)
	if err!=nil{
		fmt.Printf("insert failed(不能重复添加),err:%v\n",err)
		return
	}
	ID,err:=ret.LastInsertId()
	if err!=nil{
		fmt.Printf("get id failed,err%v\n",err)
		return
	}
	fmt.Printf("插入成功！id为%v\n",ID)

}


func main() {
	for i := 0;i < 10000000000;i++{

		insert("123456")

	}

}
