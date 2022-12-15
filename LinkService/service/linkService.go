package service

import (
	"fmt"
	"linkservice/dto"
	"linkservice/repository"
	"net/url"
)

type LinkService struct {
	linkRepository         *repository.LinkRepository
	approvedLinkRepository *repository.ApprovedLinkRepository
	userToLinkRepository   *repository.UserToLinkRepository
}

func NewLinkSevice(linkRepository *repository.LinkRepository,
	approvedLinkRepository *repository.ApprovedLinkRepository,
	userToLinkRepository *repository.UserToLinkRepository) *LinkService {
	return &LinkService{linkRepository, approvedLinkRepository, userToLinkRepository}
}

func (linkService *LinkService) CreateLink(createLinkDto *dto.CreateLinkDto, userId string) {
	u, err := url.Parse(createLinkDto.Link)
	fmt.Println(u.Host)

	approvedUrl := linkService.approvedLinkRepository.DoesLinkExist(u.Host)

	if approvedUrl != nil {
		linkFromDb := linkService.linkRepository.DoesLinkExist(createLinkDto.Link)
		if linkFromDb == nil {
			linkFromDb = linkService.linkRepository.CreateLink(approvedUrl)
			linkService.userToLinkRepository.CreateLinkToUser(userId, linkFromDb.LinkId.String())
		}
	}

	if err != nil {
		fmt.Println(err)
	}

}
