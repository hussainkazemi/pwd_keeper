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

type RandomPassword struct {
	N_letter uint8
	N_capita uint8
	N_number uint8
	N_signs  uint8
	Length   uint8
}
