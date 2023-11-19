package routes

import (
	"weekly/go/gin/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(router *gin.Engine, jobController *controllers.JobsController, applierController *controllers.ApplierController) {


	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler) )
	router.GET("/api/job", jobController.GetJobs)
	router.GET("/api/job/:jobId", jobController.GetJobById)
	router.GET("/api/job/created/:userId", jobController.GetJobByUserId)
	router.GET("/api/job/jobfilter", jobController.GetJobByCategory)
	router.GET("/api/job/applier", applierController.GetByUserId)
	router.POST("/api/job", jobController.CreateJob)
	router.POST("api/apply", applierController.ApplyJob)
	router.DELETE("/api/job/:jobId", jobController.DeleteJob)
	router.PATCH("/api/job/:jobId", jobController.UpdateJob)

}