package controllers

import (
	"net/http"
	"strconv"
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

func (controller *JobsController) CreateJob(ctx *gin.Context) {

}

func (controller *JobsController) ApplyJob(ctx *gin.Context){

}

func (controller *JobsController) UpdateJob(ctx *gin.Context) {

}

func (controller *JobsController) DeleteJob(ctx *gin.Context) {

}