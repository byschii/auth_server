package models

import (
	"time"

	"gorm.io/gorm"
)

type Resettable struct {
	ResetCode string `gorm:"type:varchar(255);unique"`
	ResetAt   time.Time
}

type User struct {
	gorm.Model
	UserName  string
	Email     string `gorm:"uniqueIndex"`
	Verified  bool
	Password  Password    `gorm:"foreignKey:UserID"`
	ApiKey    ApiKey      `gorm:"foreignKey:UserID"`
	RequstLog []RequstLog `gorm:"foreignKey:UserID"`
}

type Password struct {
	gorm.Model
	Resettable     Resettable
	UserId         uint
	HashedPassword string
	Salt           string
}

type ApiKey struct {
	gorm.Model
	Resettable Resettable
	UserID     uint
	Key        string
}

type RequstLog struct {
	gorm.Model
	UserID uint
	Method string
	Path   string
	When   time.Time
}
