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

// ApplyJob		    godoc
// @Summary			Apply job
// @Description		Api endpoint to apply job.
// @Param			input body request.ApplierRequest true "data input"
// @Produce			application/json
// @Tags			Applicants
// @Success			200 {object} response.WebResponse{}
// @Router			/job/apply [post]
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

// GetAppliedJobs	godoc
// @Summary			Get Applied Jobs
// @Description		Api endpoint to get applied job by userId.
// @Param			userId path string true "param enpoint"
// @Produce			application/json
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job/applier/{userId} [get]
func (controller *ApplierController) GetByUserId(ctx *gin.Context) {
	UserId :=  ctx.Param("userId")

	fmt.Println(UserId)
	

	jobResponse:= controller.ApplierService.FindByUserId(UserId)



	response := response.WebResponse{
		Message: "Success ambil data job berdasarkan user Id ",
		Status: "Ok",
		Data: jobResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

// CancleApply		godoc
// @Summary			Cancel Apply job
// @Description		Api Enpoint to cancel apply job.
// @Param			input body request.ApplierRequest true "data input"
// @Produce			application/json
// @Tags			Applicants
// @Success			200 {object} response.WebResponse{}
// @Router			/job/cancel [post]
func (controller *ApplierController) CancelApply(ctx *gin.Context) {

	var ApplierRequest request.ApplierRequest

	err := ctx.ShouldBindJSON(&ApplierRequest)
	
	

	errCancelApply:= controller.ApplierService.CancelApply(ApplierRequest.UserId, ApplierRequest.JobId)

	if errCancelApply != nil {
		fmt.Println(err)
		errRepsonse := response.ErrorResponse{
			Errors: err.Error(),
			Status: "Not Found",
		}
		ctx.JSON(http.StatusNotFound, errRepsonse)
		return 
	}


	webResponse := response.WebResponse{
		Message: "Berhasil Cancel Job",
		Status: "Ok",
		Data: nil,
	 }

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
