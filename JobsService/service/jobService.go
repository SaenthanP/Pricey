package service

import (
	"fmt"
	"jobservice/model"
	"jobservice/repository"
	"jobservice/constants"
	"time"
)

type JobService struct {
	jobRepository *repository.JobRepository
	ticker        *time.Ticker
}

func NewJobService(jobRepository *repository.JobRepository) *JobService {
	t := time.NewTicker(5 * time.Second)

	return &JobService{jobRepository, t}
}

func (jobService *JobService) Test() {
	fmt.Println("Test")
}

func (jobService *JobService) RetrieveJobs() {
	go func() {
		for {
			select {
			case <-jobService.ticker.C:
				jobsToRun := jobService.jobRepository.GetAllActiveJobs()
				for index := 0; index < len(jobsToRun); index++ {
					switch jobsToRun[index].JobType {
					case jobtype.SCRAPE:
						fmt.Println("test to run")
					}
					/*
					Update last run timestamp
					Delete jobs that are not repeatable
					*/
				}
			}
		}
	}()

}

func executeJobs(jobs *model.Job) {

}
