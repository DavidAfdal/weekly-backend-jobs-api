package services

import (
	"weekly/go/gin/helper"
	"weekly/go/gin/models"
	"weekly/go/gin/repository"
)

type JobsServiceImpl struct {
	JobRepo repository.JobsRepository
}

type JobsService interface {
	FindAll() []models.Job
	FindById(id int) models.Job
}

func NewJobsServiceImpl(repo repository.JobsRepository) JobsService {
	return &JobsServiceImpl{JobRepo: repo}
}

func (s *JobsServiceImpl) FindAll() []models.Job {

    result := s.JobRepo.FindAll()
	
	return result
}

func (s *JobsServiceImpl) FindById(id int) models.Job{
	result, err:= s.JobRepo.FindById(id)
	helper.ErrorPanic(err)
	return result
}