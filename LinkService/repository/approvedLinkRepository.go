package repository

import (
	"fmt"
	"linkservice/model"

	"gorm.io/gorm"
)

type ApprovedLinkRepository struct {
	db *gorm.DB
}

func NewApprovedLinkRepository(db *gorm.DB) *ApprovedLinkRepository {
	return &ApprovedLinkRepository{db}
}

func (approvedLinkRepository *ApprovedLinkRepository) DoesLinkExist(url string) *model.ApprovedLink {
	approvedLink := &model.ApprovedLink{}
	fmt.Println(url)
	approvedLinkRepository.db.First(approvedLink, "link=?", url)
	fmt.Println(approvedLink.Link)
	if approvedLink.Link == "" {
		return nil
	}
	return approvedLink
}
