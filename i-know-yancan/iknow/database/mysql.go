package database

import (
	"a/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() error {
	dsn := "root:Anyc@tcp(127.0.0.1:3306)/iknow?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.UserDatabase{}, &models.QuestionDatabase{}, &models.AnswerDatabase{}, &models.CollectDatabase{})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
