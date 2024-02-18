package main

import (
	"weekly/go/gin/config"
	"weekly/go/gin/controllers"
	_ "weekly/go/gin/docs"
	"weekly/go/gin/middleware"
	"weekly/go/gin/repository"
	"weekly/go/gin/routes"
	"weekly/go/gin/services"

	"github.com/gin-gonic/gin"
)

// @title           Gin Go Jobs Service
// @version         1.0
// @description     A Jobs management service API in Go using Gin framework.
// @contact.name   David Afdal
// @host      localhost:5000
// @BasePath  /api

func main() {
	router := gin.Default()
	router.Use(middleware.CORS())

	db := config.ConnectionDb()


	
	jobRepo := repository.NewJobsRepo(db)
	jobServis := services.NewJobsService(jobRepo)
	jobsController := controllers.NewJobsController(jobServis)
	
	applierRepo := repository.NewApplierRepo(db)
	applierServis := services.NewApplierService(applierRepo, jobRepo)
	applierController := controllers.NewApplierController(applierServis)
    

   routes.Router(router, jobsController, applierController)

	router.Run(":3000")
}