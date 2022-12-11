package repository

import (
	"gorm.io/gorm"
)

type ApprovedLinkRepository struct {
	db *gorm.DB
}

func NewApprovedLinkRepository(db *gorm.DB) *ApprovedLinkRepository {
	return &ApprovedLinkRepository{db}
}
