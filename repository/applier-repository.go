package repository

import (
	"errors"
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
	result := r.DB.Model(&models.Apllier{}).Where("user_id = ?", userId).Preload("Jobs").First(&applier)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return applier, errors.New("Applier Not Found")
	} else {
		return applier, nil
	}

}

func (r *ApllierRepo) AppendJob(job models.Job) {
	r.DB.Model(&models.Apllier{}).Association("Jobs").Append(&job)
}

func (r *ApllierRepo) DeleteApply( job models.Job, applier models.Apllier) {
	r.DB.Model(&applier).Association("Jobs").Delete(&job)
}



