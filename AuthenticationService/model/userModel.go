package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID `gorm:"primary_key; unique; type:uuid; column:UserId; default:uuid_generate_v4()"`
	Email     string
	Username  string
	LastLogin time.Time
}
