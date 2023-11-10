package services

import (
	"errors"
	"weekly/go/gin/data/request"
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
		return errors.New("Cannot Apply")
	}

   if _, err := s.ApplierRepo.FindByUserId(applier.UserId); err != nil {
	  
	  applierData := models.Apllier{
		Name: applier.Name,
		UserID: applier.UserId,
		Jobs: []models.Job{jobData},
	  }
	  s.ApplierRepo.Save(applierData)
   } else {
	
	s.ApplierRepo.AppendJob(jobData)
   }

   return nil;
}

func (s *ApplierService) FindByUserId(userId string) (models.Apllier, error) {
	data, err := s.ApplierRepo.FindByUserId(userId)
    return data,err
} 