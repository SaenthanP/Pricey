package service

import (
	"fmt"
	"linkservice/asyncmessaging"
	"linkservice/dto"
	"linkservice/model"
	"linkservice/repository"
	"log"
	"net/url"
)

type LinkService struct {
	asyncMessagingClient   *asyncmessaging.AsyncMessageClient
	linkRepository         *repository.LinkRepository
	approvedLinkRepository *repository.ApprovedLinkRepository
	userToLinkRepository   *repository.UserToLinkRepository
}

func NewLinkSevice(asyncMessagingClient *asyncmessaging.AsyncMessageClient, linkRepository *repository.LinkRepository,
	approvedLinkRepository *repository.ApprovedLinkRepository,
	userToLinkRepository *repository.UserToLinkRepository) *LinkService {
	return &LinkService{asyncMessagingClient, linkRepository, approvedLinkRepository, userToLinkRepository}
}

func (linkService *LinkService) CreateLink(createLinkDto *dto.CreateLinkDto, userId string) (*model.Link, string) {
	u, err := url.Parse(createLinkDto.Link)

	if err != nil {
		return nil, err.Error()
	}

	approvedUrl := linkService.approvedLinkRepository.DoesLinkExist(u.Host)
	if approvedUrl == nil {
		err := fmt.Sprintf("An approved Url does not exist for user provided url: %s", createLinkDto.Link)
		return nil, err
	}

	linkFromDb := linkService.linkRepository.DoesLinkExist(createLinkDto.Link)

	if linkFromDb == nil {
		linkFromDb = linkService.linkRepository.CreateLink(approvedUrl, createLinkDto.Link)
		linkService.userToLinkRepository.CreateLinkToUser(userId, linkFromDb.LinkId.String())
	} else if !linkService.userToLinkRepository.DoesLinkExistToUser(userId, linkFromDb.LinkId.String()) {
		linkService.userToLinkRepository.CreateLinkToUser(userId, linkFromDb.LinkId.String())
	}

	linkFromDb.ApprovedLink = *approvedUrl
	return linkFromDb, ""
}

func (linkService *LinkService) ScrapeJob() {
	log.Println("Scrape Job Called")
	links := linkService.linkRepository.GetAllLinks()

	for _, element := range links {
		linkService.asyncMessagingClient.CallScrape(element)
	}
}
