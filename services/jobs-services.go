package services

import (
	"fmt"
	"weekly/go/gin/data/request"
	"weekly/go/gin/data/response"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"
	"weekly/go/gin/repository"

	"github.com/go-playground/validator/v10"
)

type JobsService struct {
	JobRepo *repository.JobsRepository
	Validate *validator.Validate
}


func NewJobsServiceImpl(repo *repository.JobsRepository, validate *validator.Validate) *JobsService {
	return &JobsService{JobRepo: repo, Validate: validate}
}
 

func (s *JobsService) Create(job request.CreateJobInput) {
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
		UserId: job.UserId,
	 }

	 s.JobRepo.Save(jobModel)
}

func (s *JobsService) FindAll() []response.JobResponse {

    result := s.JobRepo.FindAll()
	fmt.Println(result)

	var jobs []response.JobResponse

	if len(result) == 0 {
		jobs = []response.JobResponse{}
	} else {
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
	}

  
	
	return jobs
}

func (s *JobsService) FindById(id int) (models.Job, error) {
	result, err := s.JobRepo.FindById(id)

	return result, err
}

func (s *JobsService) FindByUserId(userId string) []response.JobResponse {

    result := s.JobRepo.FindByUserId(userId)
	fmt.Println(result)

	var jobs []response.JobResponse
	if len(result) == 0 {
		jobs = []response.JobResponse{}
	} else {
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
	}
	
	
	return jobs
}


func (s *JobsService) FindByCategory(category string) []response.JobResponse {

    result := s.JobRepo.FindByCategory(category)
	var jobs []response.JobResponse
	if len(result) == 0 {
		jobs = []response.JobResponse{}
	} else{
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
}
	
	return jobs
}

func (s *JobsService) Update(job request.UpdateJobInput) error {
	jobData, err := s.JobRepo.FindById(job.Id)
	
	if err != nil {
		return err
	}

	jobData.Title = job.Title
	jobData.Company = job.Company
	jobData.Description = job.Description
	jobData.ImageCompany = job.ImageCompany
	jobData.Category = job.Category
	jobData.Status = job.Status
	jobData.Location = job.Location
	jobData.Salary = job.Salary
    

	s.JobRepo.Update(jobData)

	return nil
}

func (s *JobsService) Delete(jobId int) error {

	_, err :=  s.JobRepo.FindById(jobId)

	if err != nil {
		return err
	}

	s.JobRepo.Delete(jobId)

	return nil
}