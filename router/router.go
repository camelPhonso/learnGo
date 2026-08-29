package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "local.learn.go/db"
)

func FindAll(context *gin.Context) {
	movies, error := db.GetAllMovies()
	if error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return	
	}

	context.JSON(http.StatusOK, gin.H{"movies": movies})
}

func FindOne(context *gin.Context) {
	id := context.Param("movieId")
	movie, error := db.GetMovie(id)
	if error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"movie": movie})
}

func Insert(context *gin.Context) {
	var movie *db.Movie
	contextError := context.Bind(&movie)
	if contextError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": contextError.Error()})
		return
	}

	movie, error := db.CreateMovie(movie)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"movie": movie})
}

func DeleteOne(context *gin.Context) {
	id := context.Param("movieId")
	error := db.DeleteMovie(id)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Removed successfully"})
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(context *gin.Context) { context.JSON(http.StatusOK, gin.H{"message": "Server is alive"})})

	router.GET("/movies", FindAll) 
	router.GET("/movies/:movieId", FindOne)
	router.POST("/movies", Insert)
	router.DELETE("/movies/:movieId", DeleteOne)

	return router
}
