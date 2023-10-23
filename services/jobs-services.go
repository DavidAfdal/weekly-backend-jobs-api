package services

import (
	"weekly/go/gin/data/request"
	"weekly/go/gin/data/response"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"
	"weekly/go/gin/repository"

	"github.com/go-playground/validator/v10"
)

type JobsServiceImpl struct {
	JobRepo repository.JobsRepository
	Validate *validator.Validate
}

type JobsService interface {
	Create(job request.CreateJobInput)
	FindAll() []response.JobResponse
	FindById(id int) models.Job
	FindByCategory(category string) []response.JobResponse
	Update(job request.UpdateJobInput)
	Delete(jobId int)
}

func NewJobsServiceImpl(repo repository.JobsRepository, validate *validator.Validate) JobsService {
	return &JobsServiceImpl{JobRepo: repo, Validate: validate}
}


func (s *JobsServiceImpl) Create(job request.CreateJobInput) {
     err := s.Validate.Struct(job)
	 helper.ErrorPanic(err)
	 jobModel := models.Job{
		Title: job.Title,
		Description: job.Description,
		Company: job.Company,
		ImageCompany: job.ImageCompany,
		Category: job.Category,
		Status: job.Status,
		Location: job.Location,
		Salary: job.Salary,
	 }

	 s.JobRepo.Save(jobModel)
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

func (s *JobsServiceImpl) FindByCategory(category string) []response.JobResponse {

    result, err := s.JobRepo.FindByCategory(category)

	helper.ErrorPanic(err)
    
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

func (s *JobsServiceImpl) Update(job request.UpdateJobInput) {
	jobData, err := s.JobRepo.FindById(job.Id)
	helper.ErrorPanic(err)

	jobData.Title = job.Title
	jobData.Company = job.Company
	jobData.Description = job.Description
	jobData.ImageCompany = job.ImageCompany
	jobData.Category = job.Category
	jobData.Status = job.Status
	jobData.Location = job.Location
	jobData.Salary = job.Salary
    

	s.JobRepo.Update(jobData)
}

func (s *JobsServiceImpl) Delete(jobId int) {
	s.JobRepo.Delete(jobId)
}