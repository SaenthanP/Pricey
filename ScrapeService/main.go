package main

import (
	"log"
	"runtime"
	"scrapeservice/asyncmessaging"
	"scrapeservice/service"
)

/*
Maybe feature a dashbaord with live data of what is in a worker, current nubmer of workers,
*/
func main() {

	workerPool := service.NewWorkerPool(5)

	for index := 0; index < workerPool.MaxWorkers; index++ {
		workerPool.WG.Add(1)
		go service.Worker(workerPool.Jobs)
	}
	asyncmessaging.NewAsyncMessageClient(workerPool)

	runtime.Goexit()
	workerPool.WG.Wait()
	/*
		Listen to terminating command, send to a channel that it is done, have the workers listen to this then return.
		Close the channels
	*/
	//TODO create a channel, and use it to end the main go loop
	log.Println("Exit")
}
