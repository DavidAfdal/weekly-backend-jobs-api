package main

import (
	"weekly/go/gin/config"
	"weekly/go/gin/controllers"
	_ "weekly/go/gin/docs"
	"weekly/go/gin/repository"
	"weekly/go/gin/routes"
	"weekly/go/gin/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @title           Gin Go Jobs Service
// @version         1.0
// @description     A Jobs management service API in Go using Gin framework.
// @contact.name   David Afdal
// @host      localhost:5000
// @BasePath  /api

func main() {
	db := config.ConnectionDb()

	router := gin.Default()

	validate := validator.New()
    

	jobRepo := repository.NewJobsRepositoryImpl(db)
	jobServis := services.NewJobsServiceImpl(jobRepo, validate)
	jobsController := controllers.NewJobsController(jobServis)

	applierRepo := repository.NewApplierRepo(db)
	applierServis := services.NewApplierServiceImpl(applierRepo, jobRepo)
	applierController := controllers.NewApplierController(applierServis)
    
	routes.Routes(router, jobsController, applierController)

	router.Run(":5000")
}