package model

import (
	"gorm.io/gorm"
	"time"
)

type Password struct {
	gorm.Model
	Id        uint32 `gorm:"primaryKey; AUTO_INCREMENT"`
	Label     string
	Password  string
	UserId    uint32
	CreatedAt time.Time `gorm:"autoCreateTime/mil"`
	UpdatedAt time.Time `gorm:"autoUpdateTime/mil"`
}
