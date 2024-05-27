package main

import (
	"os"
	"weekly/go/gin/config"
	"weekly/go/gin/controllers"
	_ "weekly/go/gin/docs"
	"weekly/go/gin/middleware"
	"weekly/go/gin/repository"
	"weekly/go/gin/routes"
	"weekly/go/gin/services"

	dotenv "github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// @title          Workfinder Rest Api
// @version         1.0
// @description     A Jobs management service API in Go using Gin framework.
// @contact.name   David Afdal
// @host       103.157.26.181:7426
// @BasePath  /api

func main() {
	dotenv.Load(".env")
	router := gin.Default()
	router.Use(middleware.CORS())

	db := config.ConnectionDb()

    port :=  envPortOr("3000")
	
	jobRepo := repository.NewJobsRepo(db)
	jobServis := services.NewJobsService(jobRepo)
	jobsController := controllers.NewJobsController(jobServis)
	
	applierRepo := repository.NewApplierRepo(db)
	applierServis := services.NewApplierService(applierRepo, jobRepo)
	applierController := controllers.NewApplierController(applierServis)
    

   routes.Router(router, jobsController, applierController)

	router.Run(port)
}


func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
	  return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
  }