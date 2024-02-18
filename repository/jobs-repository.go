package repository

import (
	"errors"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/gorm"
)

type JobsRepository struct {
	Db *gorm.DB
}


func NewJobsRepo(Db *gorm.DB) *JobsRepository {
	return &JobsRepository{Db: Db}
}

func (r *JobsRepository) Save(job models.Job) {
	result := r.Db.Create(&job)
	helper.ErrorPanic(result.Error)
}

func (r *JobsRepository) FindAll() []models.Job {
	var jobs []models.Job

	result := r.Db.Find(&jobs)
	helper.ErrorPanic(result.Error)
	
	return jobs
}

func (r *JobsRepository) FindById(jobId int) (models.Job, error) {
	var job models.Job

	result := r.Db.Model(&models.Job{}).Where("id = ?", jobId).Preload("Aplliers").First(&job)
	
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return job, errors.New("job not found")
	} else {
		return job, nil
	}
}

func (r *JobsRepository) FindByUserId(userId string) []models.Job {
	var jobs []models.Job

	 r.Db.Where("user_id = ?", userId).First(&jobs)
	
	
	
	return jobs
}

func (r *JobsRepository) FindByCategory(category string) []models.Job {
	var jobs []models.Job

	 r.Db.Where("category = ?", category).Find(&jobs)


    return jobs
}


func (r *JobsRepository) Update(job models.Job) {
	result := r.Db.Model(&models.Job{}).Where("id = ?", job.Id).Updates(job)
	helper.ErrorPanic(result.Error)
}

func (r *JobsRepository) Delete(jobId int) {
	var job models.Job
	result := r.Db.Where("id = ?", jobId).Delete(&job)
	helper.ErrorPanic(result.Error)
}

