package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	RealName string `gorm:"size:128"`
	Avatar   string `gorm:"size:255"`
	Mobile   string `gorm:"size:128"`
	password string `gorm:"size:128;not null"`
}
