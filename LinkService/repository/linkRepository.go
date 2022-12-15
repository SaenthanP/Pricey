package repository

import (
	"linkservice/model"

	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{db}
}

func (linkRepository *LinkRepository) DoesLinkExist(url string) *model.Link {

	link := &model.Link{}
	linkRepository.db.First(link, "link=?", url)

	if link.Link == "" {
		return nil
	}
	return link
}

func (linkRepository *LinkRepository) CreateLink(approvedUrl *model.ApprovedLink) *model.Link {

	link := &model.Link{}
	link.Link = approvedUrl.Link
	link.ApprovedLinkId = approvedUrl.ApprovedLinkId

	linkRepository.db.Create(link)

	return link
}
