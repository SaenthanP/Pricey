package repository

import (
	"linkservice/model"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserToLinkRepository struct {
	db *gorm.DB
}

func NewUserToLinkRepository(db *gorm.DB) *UserToLinkRepository {
	return &UserToLinkRepository{db}
}

func (userToLinkRepository *UserToLinkRepository) DoesLinkExistToUser(userId string, linkId string) bool {

	userToLink := &model.UserToLink{}
	userToLinkRepository.db.First(userToLink, "link_id=? AND user_id=?", userId, linkId)

	return userToLink.UserId.String() != ""
}

func (userToLinkRepository *UserToLinkRepository) CreateLinkToUser(userId string, linkId string) *model.UserToLink {

	userToLink := &model.UserToLink{}
	userToLink.LinkId = uuid.MustParse(userId)
	userToLink.UserId = uuid.MustParse(linkId)
	userToLinkRepository.db.Create(userToLink)

	if userToLink.UserId.String() == "" {
		return nil
	}
	return userToLink
}
