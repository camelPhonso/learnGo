package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(AuthMiddleware())

	router.GET("/health", func(context *gin.Context) { context.JSON(http.StatusOK, gin.H{"message": "Server is alive"})})

	router.GET("/movies", FindAll) 
	router.GET("/movies/:movieId", FindOne)
	router.POST("/movies", Insert)
	router.DELETE("/movies/:movieId", DeleteOne)

	return router
}
