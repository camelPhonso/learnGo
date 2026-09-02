package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	db "local.learn.go/db"
)

func FindAll(context *gin.Context) {
	movies, error := db.GetAllMovies()
	if error != nil {
		response := FromError(error)
		context.JSON(response.Status, gin.H{"error": response.Error()})
		return	
	}

	context.JSON(http.StatusOK, gin.H{"movies": movies})
}

func FindOne(context *gin.Context) {
	id := context.Param("movieId")
	movie, error := db.GetMovie(id)
	if error != nil {
		response := FromError(error)
		context.JSON(response.Status, gin.H{"error": response.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"movie": movie})
}

func Insert(context *gin.Context) {
	var movie *db.Movie
	contextError := context.Bind(&movie)
	if contextError != nil {
		response := FromError(contextError)
		context.JSON(response.Status, gin.H{"error": response.Error()})
		return
	}

	movie, error := db.CreateMovie(movie)
	if error != nil {
		response := FromError(error)
		context.JSON(response.Status, gin.H{"error": response.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"movie": movie})
}

func DeleteOne(context *gin.Context) {
	id := context.Param("movieId")
	
	error := db.DeleteMovie(id)
	if error != nil {
		response := FromError(error)
		context.JSON(response.Status, gin.H{"error": response.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Removed successfully"})
}

func UploadFile(context *gin.Context) {
	file, err := context.FormFile("movie-list") // match the form key in the template
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movies, err := db.Read(file.Filename)
	if err != nil {
		result := FromError(err)
		context.JSON(result.Status, gin.H{"error": result.Error()})
		return
	}
	
	db.DB.Transaction( func(tx *gorm.DB) error {
		for _, movie := range movies {
			_, err := db.CreateMovie(&movie)
			if err != nil {
				return err
			}
		}
		
		return nil
	})

	context.JSON(http.StatusOK, gin.H{"file": file.Filename})
}