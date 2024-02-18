package services

import (
	"errors"
	"fmt"
	"weekly/go/gin/data/request"
	"weekly/go/gin/data/response"
	"weekly/go/gin/models"
	"weekly/go/gin/repository"
)

type ApplierService struct {
	ApplierRepo *repository.ApllierRepo 
	JobRepo *repository.JobsRepository
}


func NewApplierService(ApplierRepo *repository.ApllierRepo ,jobRepo *repository.JobsRepository) *ApplierService {
	return &ApplierService{ApplierRepo: ApplierRepo,JobRepo: jobRepo}
}

func (s *ApplierService) ApplyJob(applier request.ApplierRequest) error {
   
	jobData, err :=  s.JobRepo.FindById(applier.JobId)

	if err != nil {
        return err;
    }
	
	if applier.UserId == jobData.UserId {
		return errors.New("You Can't Apply Your Job Created")
	}

   if appliant, err := s.ApplierRepo.FindByUserId(applier.UserId); err != nil {
	  fmt.Print(err)
	  applierData := models.Apllier{
		UserID: applier.UserId,
		Jobs: []models.Job{jobData},
	  }
	  s.ApplierRepo.Save(applierData)
   } else {
	s.ApplierRepo.AppendJob(appliant, jobData)
   }

   return nil;
}

func (s *ApplierService) FindByUserId(userId string) ([]response.JobResponse) {
	var jobs []response.JobResponse
	data, _ := s.ApplierRepo.FindByUserId(userId)
	if len(data.Jobs) == 0 {
		jobs = []response.JobResponse{}
	} else {
		for _, value := range data.Jobs {
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

func (s *ApplierService) CancelApply(userId string, jobId int) error{
   dataApplier, err := s.ApplierRepo.FindByUserId(userId)

   if err != nil {
	return err
   }

   dataJob, err := s.JobRepo.FindById(jobId)

   if err != nil {
	return err
   }


s.ApplierRepo.DeleteApply(dataJob, dataApplier)

   return nil
}