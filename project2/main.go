package main

import (
	"fmt"
	"project2/controllers"
	"project2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server.......")

	masterDB := database.Postgres()
	controllers := controllers.New(masterDB)

	app := gin.Default()

	app.POST("/books/", controllers.CreateBook)
	app.GET("/books/:id", controllers.GetBook)
	app.GET("/books/", controllers.GetAllBooks)
	app.PUT("/books/:id", controllers.UpdateBook)
	app.DELETE("/books/:id", controllers.DeleteBook)
	app.Run(":8080")
}
