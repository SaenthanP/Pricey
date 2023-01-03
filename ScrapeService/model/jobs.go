package model

type Job struct{
	JobType string
	Executor func() 
}