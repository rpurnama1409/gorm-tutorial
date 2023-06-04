package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama   string `gorm:"not null"`
	Alamat string `gorm:"not null"`
}

type Users []User
