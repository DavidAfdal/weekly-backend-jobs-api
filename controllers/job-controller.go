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
	jobService *services.JobsService
}

func NewJobsController(service *services.JobsService) *JobsController {
	return &JobsController{
		jobService: service,
	}
}

// GetJobs			godoc
// @Summary		    Get All Jobs.
// @Description		Api Enpoint to Get All Data Jobs
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job [get]
func (controller *JobsController) GetJobs(ctx *gin.Context) {
  jobResponse := controller.jobService.FindAll()

  response := response.WebResponse{
	 Message: "Sucess Get All Data",
	 Status: "OK",
	 Data: jobResponse,
  }

  ctx.Header("Content-Type", "application/json")
  ctx.JSON(http.StatusOK, response)
}


// GetJobById		godoc
// @Summary			Get Single Job By Id.
// @Param			jobId path int true "Param Endpoint"
// @Description		Api Enpoint to Get Single Job By Id
// @Produce			application/json
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job/{jobId} [get]
func (controller *JobsController) GetJobById(ctx *gin.Context) {
	jobId :=  ctx.Param("jobId")
	id, err := strconv.Atoi(jobId)
	helper.ErrorPanic(err)

	jobResponse, err := controller.jobService.FindById(id)



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
		Message: "Succes Get Data By Id",
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
		Message: "Success Get Data Job By " + jobCategory,
		Status: "Ok",
		Data: jobResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}


// GetSharedJobs	godoc
// @Summary			Get Shared jobs by userId.
// @Param			userId path string true "Param Endpoint"
// @Description		Api Endpoint To Get Shared Jobs.
// @Produce			application/json
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job/created/{userId} [get]
func (controller *JobsController) GetJobByUserId(ctx *gin.Context) {

	UserId :=  ctx.Param("userId")
	
	fmt.Println(UserId)

	jobResponse := controller.jobService.FindByUserId(UserId)


	response := response.WebResponse{
		Message: "Success Get Data Job By UserId",
		Status: "Ok",
		Data: jobResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

// CreateJob		godoc
// @Summary			Create New Job
// @Description		Api Endpoint To Create New Job.
// @Param			input body request.CreateJobInput true "Data Input"
// @Produce			application/json
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job [post]
func (controller *JobsController) CreateJob(ctx *gin.Context) {
     var JobRequest request.CreateJobInput

	 err := ctx.ShouldBindJSON(&JobRequest)

	 if errBindResult := helper.BindErrorHandler(err); errBindResult != nil {
		errorResponse := response.ErrorResponse{
			Status: "Bad Request",
			Errors: errBindResult,
		}
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	 controller.jobService.Create(JobRequest)

     webResponse := response.WebResponse{
		Message: "Succes add data job",
		Status: "Created",
		Data: nil,
	 }

	 ctx.Header("Content-Type", "application/json")
	 ctx.JSON(http.StatusCreated, webResponse)
}


// UpdateJob		godoc
// @Summary			Update Job
// @Description		Api Endpoint To Update Job
// @Param			jobId path int true "Param Enpoint"
// @Param			input body request.UpdateJobInput true "Data Input"
// @Produce			application/json
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job/{jobId} [patch]
func (controller *JobsController) UpdateJob(ctx *gin.Context) {

	updateJobRequest := request.UpdateJobInput{}
	err := ctx.ShouldBindJSON(&updateJobRequest)

	if errBindResult := helper.BindErrorHandler(err); errBindResult != nil {
		errorResponse := response.ErrorResponse{
			Status: "Bad Request",
			Errors: errBindResult,
		}
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

    jobId := ctx.Param("jobId")
	id, err := strconv.Atoi(jobId)
	helper.ErrorPanic(err)

	if err:= controller.jobService.Update(updateJobRequest, id); err != nil {
		errorResponse := response.ErrorResponse{
			Status: "Not Found",
			Errors: err.Error(),
		}
		ctx.JSON(http.StatusNotFound, errorResponse)
		return
	}

	webResponse:= response.WebResponse {
		Status: "Ok",
		Message: "Success Update Data Job",
		Data: nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// DeleteJob		godoc
// @Summary			Delete job
// @Description		Api Endpoint for Delete job.
// @Param			jobId path int true "param endpoint"
// @Produce			application/json
// @Tags			Jobs
// @Success			200 {object} response.WebResponse{}
// @Router			/job/{jobId} [delete]
func (controller *JobsController) DeleteJob(ctx *gin.Context) {
	jobId := ctx.Param("jobId")
	id, err := strconv.Atoi(jobId)
	helper.ErrorPanic(err)

	if err:= controller.jobService.Delete(id); err != nil {
		errorResponse := response.ErrorResponse{
			Status: "Not Found",
			Errors: err.Error(),
		}
		ctx.JSON(http.StatusNotFound, errorResponse)
		return
	}

	webResponse := response.WebResponse{
		Status: "Ok",
		Message: "Sucses Delete Data Job",
		Data: nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}