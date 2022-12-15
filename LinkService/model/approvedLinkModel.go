package model

import (
	"github.com/google/uuid"
)

type ApprovedLink struct {
	ApprovedLinkId uuid.UUID `gorm:"primary_key; unique; type:uuid; column:ApprovedLinkId; default:uuid_generate_v4()"`
	HtmlToScrape   string
	Link           string
}
