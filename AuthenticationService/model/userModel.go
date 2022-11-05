package model

import (
	"time"
)

type User struct {
	UserId    string
	Email     string
	Username  string
	LastLogin time.Time
}
