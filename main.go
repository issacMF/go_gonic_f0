package main

import (
	"learn_gin/controllers"
	"learn_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.Index)
	r.POST("/book", controllers.New)
	r.GET("/book/:id", controllers.Show)
	r.PATCH("/book/:id", controllers.Update)
	r.DELETE("/book/:id", controllers.Delete)

	r.Run()
}
