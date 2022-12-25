package repository

import (
	"fmt"
	"jobservice/model"
	"time"

	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db}
}

func (jobRepository *JobRepository) GetAllActiveJobs() []model.Job{
	jobs := []model.Job{}
	currTime := time.Now()
	jobRepository.db.Find(&jobs, "to_run_at < ? AND EXTRACT(DAY from (?-last_run))>1", currTime, currTime)
	fmt.Println(jobs)
	return jobs
}
