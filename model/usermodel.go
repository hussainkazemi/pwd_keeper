package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id          uint32 `gorm:"primaryKey; AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(100); not null"`
	PhoneNumber string `gorm:"type:varchar(15); not null"`
	UserName    string
	Password    string
	Passwords   []Password `gorm:"foreignKey:UserId;references:Id"`
	CreatedAt   time.Time  `gorm:"autoCreateTime/mil"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime/mil"`
}

func UserNameIsValid(userName string) bool {
	return len(userName) > 3
}
