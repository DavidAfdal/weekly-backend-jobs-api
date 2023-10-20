package services

import (
	"weekly/go/gin/data/response"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"
	"weekly/go/gin/repository"
)

type JobsServiceImpl struct {
	JobRepo repository.JobsRepository
}

type JobsService interface {
	FindAll() []response.JobResponse
	FindById(id int) models.Job
}

func NewJobsServiceImpl(repo repository.JobsRepository) JobsService {
	return &JobsServiceImpl{JobRepo: repo}
}

func (s *JobsServiceImpl) FindAll() []response.JobResponse {

    result := s.JobRepo.FindAll()
    
	var jobs []response.JobResponse
	for _, value := range result {
		job := response.JobResponse{
			Id: value.Id,
			Title: value.Title,
			Company: value.Company,
			ImageCompany: value.ImageCompany,
			Status: value.Status,
			Category: value.Category,
			CreatedAt: value.CreatedAt,
		}
		jobs = append(jobs, job)
	}
	
	return jobs
}

func (s *JobsServiceImpl) FindById(id int) models.Job{
	result, err:= s.JobRepo.FindById(id)
	helper.ErrorPanic(err)
	return result
}