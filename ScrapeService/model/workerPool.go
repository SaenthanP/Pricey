package model

import "sync"

type WorkerPool struct {
	MaxWorkers int
	Jobs       chan Job
	WG         *sync.WaitGroup
	//Results???

}
