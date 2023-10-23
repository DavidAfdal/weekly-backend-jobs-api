package main

import (
	"os"
	"weekly/go/gin/config"
	"weekly/go/gin/controllers"
	"weekly/go/gin/repository"
	"weekly/go/gin/routes"
	"weekly/go/gin/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func main() {
	db := config.ConnectionDb()

	router := gin.Default()

	validate := validator.New()
    

	jobRepo := repository.NewJobsRepositoryImpl(db)
	jobServis := services.NewJobsServiceImpl(jobRepo, validate)
	jobsController := controllers.NewJobsController(jobServis)
    
	routes.JobRoutes(router, jobsController)

	router.Run(":" + os.Getenv("PORT"))
}