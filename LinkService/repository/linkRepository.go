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

func (linkRepository *LinkRepository) CreateLink(approvedUrl *model.ApprovedLink, url string) *model.Link {

	link := &model.Link{}
	link.Link = url
	link.ApprovedLinkId = approvedUrl.ApprovedLinkId

	linkRepository.db.Create(link)

	return link
}

func (linkRepository *LinkRepository) GetAllLinks() []model.Link {
	links := []model.Link{}
	linkRepository.db.Preload("ApprovedLink").Find(&links)

	return links
}
