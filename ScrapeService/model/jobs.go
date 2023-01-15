package model

type Job struct {
	JobType  string
	Executor Executor
	MetaData []byte
}
