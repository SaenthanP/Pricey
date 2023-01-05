package service

import (
	"log"
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
			log.Println("Worker called")
			time.Sleep(10 * time.Second)

		}
	}
}
