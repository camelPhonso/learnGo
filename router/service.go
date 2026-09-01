package router

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
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

	content, err := os.Open(file.Filename)
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		var newMovie db.Movie

		line := []byte(fmt.Sprintf("{%s}", scanner.Text()))
		err := json.Unmarshal(line, &newMovie)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		db.CreateMovie(&newMovie)
	}

	context.JSON(http.StatusOK, gin.H{"file": file.Filename})
}