package models

import "gorm.io/gorm"

type UserDatabase struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;"`
	Password string `gorm:"varchar(25);not null;"`
	Mail     string `gorm:"varchar(25);not null;"`
}
