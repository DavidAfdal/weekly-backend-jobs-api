package routes

import (
	"weekly/go/gin/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(router *gin.Engine, jobController *controllers.JobsController, applierController *controllers.ApplierController) {
    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler) )
	route := router.Group("/api")
	
	route.POST("/apply", applierController.ApplyJob)

	jobRouter := router.Group("/api/job")
	jobRouter.GET("/", jobController.GetJobs)
	jobRouter.GET("/:jobId", jobController.GetJobById)
	jobRouter.GET("/created/:userId", jobController.GetJobByUserId)
	jobRouter.GET("/jobfilter", jobController.GetJobByCategory)
	jobRouter.GET("/applier", applierController.GetByUserId)
	jobRouter.POST("/", jobController.CreateJob)
	jobRouter.DELETE("/:jobId", jobController.DeleteJob)
	jobRouter.PATCH("/:jobId", jobController.UpdateJob)

}