package main

import (
	"fmt"
	"runtime"
	"scrapeservice/model"
	"scrapeservice/service"
	"time"
)

func main() {

	workerPool := service.NewWorkerPool(2)

	for index := 0; index < workerPool.MaxWorkers; index++ {
		workerPool.WG.Add(1)
		go service.Worker(workerPool.Jobs)
	}

	fmt.Println("testing 123")

	// ctx, cancel := context.WithCancel(context.TODO())
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("okay?")
	// 			return
	// 		}
	// 	}
	// }()

	time.Sleep(5 * time.Second)

	job1 := model.Job{JobType: "scrape", Executor: nil}

	workerPool.Jobs <- job1
	job2 := model.Job{JobType: "scrape", Executor: nil}

	job3 := model.Job{JobType: "scrape", Executor: nil}

	job4 := model.Job{JobType: "scrape", Executor: nil}
	job5 := model.Job{JobType: "scrape", Executor: nil}
	job6 := model.Job{JobType: "scrape", Executor: nil}
	job7 := model.Job{JobType: "scrape", Executor: nil}
	workerPool.Jobs <- job2
	fmt.Println("Put in job2")
	workerPool.Jobs <- job3
	fmt.Println("Put in job3")

	workerPool.Jobs <- job4
	fmt.Println("Put in job4")
	workerPool.Jobs <- job5
	fmt.Println("Put in job5")

	workerPool.Jobs <- job6
	fmt.Println("Put in job6")

	workerPool.Jobs <- job7
	fmt.Println("Put in job7")
	// cancel()
	// time.Sleep(10 * time.Second)
	for {

	}
	runtime.Goexit()
	workerPool.WG.Wait()

	//TODO create a channel, and use it to end the main go loop
	fmt.Println("Exit")
}
