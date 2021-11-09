package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	ApiKey       ApiKey      `gorm:"foreignKey:UserID"`
	RequstLog    []RequstLog `gorm:"foreignKey:UserID"`
}

type ApiKey struct {
	gorm.Model
	UserID    uint
	Key       string
	CodeReset string
	Resetting bool
}

type RequstLog struct {
	gorm.Model
	UserID uint
	Method string
	Path   string
}
