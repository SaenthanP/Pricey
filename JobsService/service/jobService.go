package service

import (
	"fmt"
	"jobservice/repository"
)

type JobService struct {
	jobRepository *repository.JobRepository
}

func NewJobService(jobRepository *repository.JobRepository) *JobService {
	return &JobService{jobRepository}
}

func (jobService *JobService) Test() {
	fmt.Println("Test")
}
