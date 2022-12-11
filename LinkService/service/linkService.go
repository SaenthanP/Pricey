package service

import (
	"linkservice/dto"
	"linkservice/repository"
)

type LinkService struct {
	linkRepository *repository.LinkRepository
	approvedLinkRepository *repository.ApprovedLinkRepository
	userToLinkRepository *repository.UserToLinkRepository
}

func NewLinkSevice(linkRepository *repository.LinkRepository,
				   approvedLinkRepository *repository.ApprovedLinkRepository,
				   userToLinkRepository *repository.UserToLinkRepository) *LinkService {

	return &LinkService{linkRepository,approvedLinkRepository,userToLinkRepository }
}

func(linkService *LinkService)CreateLink(createLinkDto *dto.CreateLinkDto){
	
}