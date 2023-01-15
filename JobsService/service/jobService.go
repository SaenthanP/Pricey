package service

import (
	"context"
	"fmt"
	jobtype "jobservice/constants"
	"jobservice/model"
	"jobservice/proto"
	"jobservice/repository"
	"log"
	"time"
)

type JobService struct {
	protoClient   proto.CallScrapeServiceClient
	jobRepository *repository.JobRepository
	ticker        *time.Ticker
}

func NewJobService(protoClient proto.CallScrapeServiceClient, jobRepository *repository.JobRepository) *JobService {
	t := time.NewTicker(5 * time.Second)

	return &JobService{protoClient, jobRepository, t}
}

func (jobService *JobService) Test() {
	fmt.Println("Test")
}

func (jobService *JobService) RetrieveJobs() {
	req := &proto.Request{}

	jobService.protoClient.CallScrape(context.Background(), req)
	go func() {
		for {
			select {
			case <-jobService.ticker.C:
				jobsToRun := jobService.jobRepository.GetAllActiveJobs()
				for index := 0; index < len(jobsToRun); index++ {
					switch jobsToRun[index].JobType {
					case jobtype.SCRAPE:
						req := &proto.Request{}

						if _, err := jobService.protoClient.CallScrape(context.Background(), req); err == nil {
							jobService.jobRepository.UpdateLastRunByJob(jobsToRun[index])
						} else {
							log.Fatalf("Something went wrong with executing scrape job: %v", err)
						}
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
