package models

import "gorm.io/gorm"

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
type CollectDatabase struct {
	UserID     int `gorm:"type:int;not null;"`
	QuestionID int `gorm:"type:int;not null;"`
}
