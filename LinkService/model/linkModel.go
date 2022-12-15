package model

import (
	"github.com/google/uuid"
)

type Link struct {
	LinkId         uuid.UUID `gorm:"primary_key; unique; type:uuid; column:LinkId; default:uuid_generate_v4()"`
	ApprovedLinkId uuid.UUID
	ApprovedLink   ApprovedLink `gorm:"references:ApprovedLinkId"`
	Link 			string
}
