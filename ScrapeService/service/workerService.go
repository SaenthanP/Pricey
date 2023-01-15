package service

import (
	"encoding/json"
	"scrapeservice/model"
	"sync"
	"time"
)

func NewWorkerPool(maxWorkers int) *model.WorkerPool {
	jobs := make(chan model.Job, maxWorkers)
	wg := &sync.WaitGroup{}
	return &model.WorkerPool{MaxWorkers: maxWorkers, Jobs: jobs, WG: wg}
}

func Worker(jobs chan model.Job) {
	for {
		select {
		case job := <-jobs:
			job.Executor(job.MetaData, job.JobType)
			time.Sleep(10 * time.Second)

		}
	}
}

var Exec = model.Executor(func(data []byte, jobType string) {
	if jobType == "scrape" {
		scrapeLink(data)
	}
})

func scrapeLink(data []byte) {
	link := model.Link{}
	json.Unmarshal(data, &link)
}
