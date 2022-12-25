package main

import (
	"fmt"
	"jobservice/database"
	"jobservice/repository"
	"jobservice/service"
	"runtime"
)

func main() {
	db := database.SetDB()
	jobRepository := repository.NewJobRepository(db)

	jobService := service.NewJobService(jobRepository)
	jobService.Test()

	jobService.RetrieveJobs()
	fmt.Println("test")
	runtime.Goexit()
	//TODO create a channel, and use it to end the main go loop
	fmt.Println("Exit")

}
