package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"weekly/go/gin/data/request"
	"weekly/go/gin/data/response"
	"weekly/go/gin/helper"
	"weekly/go/gin/services"

	"github.com/gin-gonic/gin"
)


type JobsController struct {
	jobService services.JobsService
}

func NewJobsController(service services.JobsService) *JobsController {
	return &JobsController{
		jobService: service,
	}
}


func (controller *JobsController) GetJobs(ctx *gin.Context) {
  jobResponse := controller.jobService.FindAll()

  response := response.WebResponse{
	 Message: "Berhasil Fetch List Data Jobs",
	 Status: "OK",
	 Data: jobResponse,
  }

  ctx.Header("Content-Type", "application/json")
  ctx.JSON(http.StatusOK, response)
}

func (controller *JobsController) GetJobById(ctx *gin.Context) {
	jobId :=  ctx.Param("jobId")
	id, err := strconv.Atoi(jobId)
	helper.ErrorPanic(err)

	jobResponse := controller.jobService.FindById(id)
	response := response.WebResponse{
		Message: "Succes Ambil Data Dari Id",
		Status: "Ok",
		Data: jobResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}


func (controller *JobsController) GetJobByCategory(ctx *gin.Context) {
	jobCategory := ctx.Query("category")
	
	fmt.Println(jobCategory)

	jobResponse := controller.jobService.FindByCategory(jobCategory)

	response := response.WebResponse{
		Message: "Success ambil data job berdasarkan catregory " + jobCategory,
		Status: "Ok",
		Data: jobResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func (controller *JobsController) CreateJob(ctx *gin.Context) {
     var JobRequest request.CreateJobInput

	 err := ctx.ShouldBindJSON(&JobRequest)
	 helper.ErrorPanic(err)

	 controller.jobService.Create(JobRequest)

     webResponse := response.WebResponse{
		Message: "Berhasil menambahkan data job",
		Status: "Ok",
		Data: nil,
	 }

	 ctx.Header("Content-Type", "application/json")
	 ctx.JSON(http.StatusOK, webResponse)
}

func (controller *JobsController) ApplyJob(ctx *gin.Context){

}

func (controller *JobsController) UpdateJob(ctx *gin.Context) {


	updateJobRequest := request.UpdateJobInput{}
	err := ctx.ShouldBindJSON(&updateJobRequest)
	helper.ErrorPanic(err)
	
    jobId := ctx.Param("jobId")
	id, err := strconv.Atoi(jobId)
	helper.ErrorPanic(err)
	updateJobRequest.Id = id

	controller.jobService.Update(updateJobRequest)

	webResponse:= response.WebResponse {
		Status: "Ok",
		Message: "Berhasil Update Job",
		Data: nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *JobsController) DeleteJob(ctx *gin.Context) {
	jobId := ctx.Param("jobId")
	id, err := strconv.Atoi(jobId)
	helper.ErrorPanic(err)

	controller.jobService.Delete(id)

	webResponse := response.WebResponse{
		Status: "Ok",
		Message: "Berhasil menghapus data Job",
		Data: nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}