package main

import (
	"os"
	"weekly/go/gin/config"
	"weekly/go/gin/controllers"
	"weekly/go/gin/repository"
	"weekly/go/gin/routes"
	"weekly/go/gin/services"

	"github.com/gin-gonic/gin"
)


func main() {
	db := config.ConnectionDb()

	router := gin.Default()
    

	jobRepo := repository.NewJobsRepositoryImpl(db)
	jobServis := services.NewJobsServiceImpl(jobRepo)
	jobsController := controllers.NewJobsController(jobServis)
    
	routes.JobRoutes(router, jobsController)

	router.Run(":" + os.Getenv("PORT"))
}