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

// FindAllJobs 		godoc
// @Summary			Get All Jobs.
// @Description		Return list of jobs.
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


// FindJobsById		godoc
// @Summary			Get Single jobs by id.
// @Param			jobId path string true "get jobs by id"
// @Description		Return the job whoes jobId value matches id
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

// CreateTags		godoc
// @Summary			Create job
// @Description		Save job data in Db.
// @Param			tags body request.CreateJobInput true "Create job"
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
	updateJobRequest.Id = id

	if err:= controller.jobService.Update(updateJobRequest); err != nil {
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