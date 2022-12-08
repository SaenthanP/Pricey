package service

import (
	"scraperservice/repository"
)

type LinkService struct {
	linkRepository *repository.LinkRepository
}

func NewLinkSevice(linkRepository *repository.LinkRepository) *LinkService {
	return &LinkService{linkRepository}
}
