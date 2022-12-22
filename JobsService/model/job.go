package model

import (
	"time"

	"github.com/google/uuid"
)


type Job struct{
	JobId uuid.UUID `gorm:"primary_key; unique; type:uuid; column:ApprovedLinkId; default:uuid_generate_v4()"`
	JobType	string
	IsReoccuring bool
	LastRun	time.Time
	ToRunAt time.Time
}