package model

import (
	"time"
)

type User struct {
	UserId    uint `gorm:"primaryKey"`
	Email     string
	Username  string
	LastLogin time.Time
}
