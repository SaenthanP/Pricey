package repository

import (
	"gorm.io/gorm"
)

type UserToLinkRepository struct {
	db *gorm.DB
}

func NewUserToLinkRepository(db *gorm.DB) *UserToLinkRepository {
	return &UserToLinkRepository{db}
}
