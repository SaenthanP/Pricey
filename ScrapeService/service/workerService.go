package service

import (
	"fmt"
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
		case <-jobs:
			fmt.Println("works????????")
			time.Sleep(10 * time.Second)

		}
	}
}
