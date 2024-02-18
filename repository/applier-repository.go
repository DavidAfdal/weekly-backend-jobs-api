package repository

import (
	"errors"
	"log"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/gorm"
)

type ApllierRepo struct {
	DB *gorm.DB
}


func NewApplierRepo(db *gorm.DB) *ApllierRepo {
	return &ApllierRepo{DB: db}
}


func (r *ApllierRepo) Save(applier models.Apllier){
	result := r.DB.Create(&applier)
	helper.ErrorPanic(result.Error)
}

func (r *ApllierRepo) FindByUserId(userId string) (models.Apllier, error) {
	var applier models.Apllier
	result := r.DB.Where("user_id = ?", userId).Preload("Jobs").First(&applier)
	log.Print(applier.Jobs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	   return applier, errors.New("applier not found")
	} else {
		return applier, nil
	}

}

func (r *ApllierRepo) AppendJob(appliant models.Apllier, job models.Job) {
	r.DB.Model(&appliant).Association("Jobs").Append(&job)
}

func (r *ApllierRepo) DeleteApply( job models.Job, applier models.Apllier) {
	r.DB.Transaction(func(tx *gorm.DB) error {


		tx.Model(&applier).Association("Jobs").Delete(&job)
		if(len(applier.Jobs) == 0) {
			tx.Delete(&applier)
		}
		return nil
	})
}



