package models

import "gorm.io/gorm"

type QuestionsParam struct {
	QUE      string `form:"que" json:"que" binding:"required"`
	UserName string `form:"username" json:"username" binding:"required"`
}

type AnswersParam struct {
	ANS        string `form:"ans" json:"ans" binding:"required"`
	UserName   string `form:"UserName" json:"UserName" binding:"required"`
	QuestionID int
}
type QuestionDatabase struct {
	gorm.Model
	Great    int    `gorm:"type:int;not null;"`
	UserName string `gorm:"varchar(20);not null;"`
	QUE      string `gorm:"varchar(2000);not null;"`
	AnsNum   int    `gorm:"type:int;not null;"`
}
type AnswerDatabase struct {
	gorm.Model
	ANS        string `gorm:"varchar(2000);not null;"`
	UserName   string `gorm:"varchar(20);not null;"`
	QuestionID int    `gorm:"type:int;not null;"`
}
type QuestionHot struct {
	Great     int
	UserParam string
	QUE       int
	AnsNum    int
}
type Like struct {
	Num  int
}
