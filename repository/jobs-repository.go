package repository

import (
	"errors"
	"weekly/go/gin/data/request"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/gorm"
)


type JobsRepository interface {
	Save(job models.Job)
	FindAll() []models.Job
	FindById(jobId int) (job models.Job, err error)
	FindByCategory(category string) (jobs []models.Job, err error)
	Update(job models.Job) 
	Delete(jobId int)
}


type JobsRepositoryImpl struct {
	Db *gorm.DB
}


func NewJobsRepositoryImpl(Db *gorm.DB) JobsRepository {
	return &JobsRepositoryImpl{Db: Db}
}

func (r *JobsRepositoryImpl) Save(job models.Job) {
	result := r.Db.Create(&job)
	helper.ErrorPanic(result.Error)
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

func (r *JobsRepositoryImpl) FindByCategory(category string) (job []models.Job, err error) {
	var jobs []models.Job

	result := r.Db.Where("category = ?", category).Find(&jobs)
	helper.ErrorPanic(result.Error)
	return jobs, err
}


func (r *JobsRepositoryImpl) Update(job models.Job) {
	var updateJob = request.UpdateJobInput{
		Id: job.Id,
		Title: job.Title,
		Description: job.Description,
		Company: job.Company,
		ImageCompany: job.ImageCompany,
		Category: job.Category,
		Status: job.Status,
		Salary: job.Salary,
	}

	result := r.Db.Model(&job).Updates(updateJob)
	helper.ErrorPanic(result.Error)
}

func (r *JobsRepositoryImpl) Delete(jobId int) {
	var job models.Job
	result := r.Db.Where("id = ?", jobId).Delete(&job)
	helper.ErrorPanic(result.Error)
}

