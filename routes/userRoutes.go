package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup ) {
	userRoute := router.Group("/user")

	userRoute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server",
		})
	})

} 