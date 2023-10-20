package repository

import (
	"errors"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/gorm"
)


type JobsRepository interface {
	FindAll() []models.Job
	FindById(id int) (job models.Job, err error)
}


type JobsRepositoryImpl struct {
	Db *gorm.DB
}


func NewJobsRepositoryImpl(Db *gorm.DB) JobsRepository {
	return &JobsRepositoryImpl{Db: Db}
}


func (r *JobsRepositoryImpl) FindAll() []models.Job {
	var jobs []models.Job

	result := r.Db.Find(&jobs)
	helper.ErrorPanic(result.Error)
	return jobs
}

func (r *JobsRepositoryImpl) FindById(jobId int) (jobs models.Job, err error) {
	var job models.Job

	result := r.Db.Find(&job, jobId)
	
	if result != nil {
       return job, nil
	} else {
		return job, errors.New("tag is not found")
	}
}


