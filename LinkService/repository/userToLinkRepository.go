package repository

import (
	"fmt"
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
	userToLinkRepository.db.First(userToLink, "link_id=? AND user_id=?",linkId, userId )
	fmt.Println(linkId)
	fmt.Println(userId)
	fmt.Println(userToLink.UserId.String())
	return userToLink.UserId!= uuid.Nil
}

func (userToLinkRepository *UserToLinkRepository) CreateLinkToUser(userId string, linkId string) *model.UserToLink {

	userToLink := &model.UserToLink{}
	userToLink.LinkId = uuid.MustParse(linkId)
	userToLink.UserId = uuid.MustParse(userId)
	userToLinkRepository.db.Create(userToLink)

	if userToLink.UserId==uuid.Nil{
		return nil
	}
	return userToLink
}
