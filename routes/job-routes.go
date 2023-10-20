package routes

import (
	"weekly/go/gin/controllers"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine, jobController *controllers.JobsController) {

	jobRouter := router.Group("/api")
	jobRouter.GET("/job", jobController.GetJobs)
	jobRouter.GET("/job/:jobId", jobController.GetJobById)
	jobRouter.POST("/job", jobController.CreateJob)
}