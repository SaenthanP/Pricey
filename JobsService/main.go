package main

import (
	"fmt"
	"jobservice/database"
	"jobservice/repository"
	"jobservice/service"
	"runtime"
	"time"
)

func main() {
	db:= database.SetDB()
	jobRepository:=repository.NewJobRepository(db)

	jobService:=service.NewJobService(jobRepository)
	jobService.Test()

	// go func() {
	// 	fmt.Println("testignt")
	// }()
	// fmt.Println("testing3gggg")
	// time.Sleep(500)
	t := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-t.C:
				fmt.Println("running?")
			}
		}
	}()
	fmt.Println("test")
	runtime.Goexit()
	//TODO create a channel, and use it to end the main go loop 
	fmt.Println("Exit")



}
