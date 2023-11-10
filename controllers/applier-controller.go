package controllers

import (
	"fmt"
	"net/http"
	"weekly/go/gin/data/request"
	"weekly/go/gin/data/response"
	"weekly/go/gin/helper"
	"weekly/go/gin/services"

	"github.com/gin-gonic/gin"
)

type ApplierController struct {
	ApplierService *services.ApplierService
}

func NewApplierController(ApplierService *services.ApplierService) *ApplierController {
	return &ApplierController{
		ApplierService: ApplierService,
	}
}


func (controller *ApplierController) ApplyJob(ctx *gin.Context) {
	var ApplierRequest request.ApplierRequest

	err := ctx.ShouldBindJSON(&ApplierRequest)

	if errBindResult := helper.BindErrorHandler(err); errBindResult != nil {
	   errorResponse := response.ErrorResponse{
		   Status: "Bad Request",
		   Errors: errBindResult,
	   }
	   ctx.JSON(http.StatusBadRequest, errorResponse)
	   return
   }

   errorApply := controller.ApplierService.ApplyJob(ApplierRequest)

	fmt.Println(errorApply)
   

   if errorApply != nil {
	errRepsonse := response.ErrorResponse{
		Errors: errorApply.Error(),
		Status: "Bad Request",
	}
	ctx.JSON(http.StatusBadRequest, errRepsonse)
	return 
}

   webResponse := response.WebResponse{
	Message: "Berhasil Apply Job",
	Status: "Ok",
	Data: nil,
 }

 ctx.Header("Content-Type", "application/json")
 ctx.JSON(http.StatusOK, webResponse)



}

func (controller *ApplierController) GetByUserId(ctx *gin.Context) {
	UserId :=  ctx.Query("userId")

	fmt.Println(UserId)
	

	jobResponse, err:= controller.ApplierService.FindByUserId(UserId)

	if err != nil {
		fmt.Println(err)
		errRepsonse := response.ErrorResponse{
			Errors: err.Error(),
			Status: "Not Found",
		}
		ctx.JSON(http.StatusNotFound, errRepsonse)
		return 
	}


	response := response.WebResponse{
		Message: "Success ambil data job berdasarkan user Id ",
		Status: "Ok",
		Data: jobResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}