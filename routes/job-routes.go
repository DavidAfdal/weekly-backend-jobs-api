package routes

import (
	"weekly/go/gin/controllers"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine, jobController *controllers.JobsController) {

	jobRouter := router.Group("/api")
	jobRouter.GET("/job", jobController.GetJobs)
	jobRouter.GET("/job/:jobId", jobController.GetJobById)
	jobRouter.GET("/jobfilter", jobController.GetJobByCategory)
	jobRouter.POST("/job", jobController.CreateJob)
	jobRouter.DELETE("/job/:jobId", jobController.DeleteJob)
	jobRouter.PATCH("/job/:jobId", jobController.UpdateJob)
}