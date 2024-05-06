package services

import (
	"fmt"
	"weekly/go/gin/data/request"
	"weekly/go/gin/data/response"
	"weekly/go/gin/models"
	"weekly/go/gin/repository"
)

type JobsService struct {
	JobRepo *repository.JobsRepository
}


func NewJobsService(repo *repository.JobsRepository) *JobsService {
	return &JobsService{JobRepo: repo}
}
 

func (s *JobsService) Create(job request.CreateJobInput) {

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

func (s *JobsService) FindById(id int) (response.JobResponse , error) {
	
	result, err := s.JobRepo.FindById(id)
	appliers := make([]string, len(result.Aplliers))

	for i, applier := range result.Aplliers {
		appliers[i] = applier.UserID
	}

	job := response.JobResponse{
		Id: result.Id,
		Title: result.Title,
		Description: result.Description,
		Company: result.Company,
		ImageCompany: result.ImageCompany,
		Location: result.Location,
		Salary: result.Salary,
		UserId: result.UserId,
		Status: result.Status,
		Category: result.Category,
		CreatedAt: result.CreatedAt,
		Appliers : appliers,
	}
		return job, err

	
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

func (s *JobsService) Update(job request.UpdateJobInput, id int) error {
	jobData, err := s.JobRepo.FindById(id)
	
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