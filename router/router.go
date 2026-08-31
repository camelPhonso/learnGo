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
	
	// html blob is served to public pages 
	// authmiddleware is not applied to this group
	router.LoadHTMLGlob("templates/*")
	public := router.Group("/public")
	{
		public.GET("/", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", gin.H{"title": "Uploading a file!"})
		})
	}
	
	return router
}
