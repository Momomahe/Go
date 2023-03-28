package main

import (
	"testing/controller"
	"testing/database"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	db := database.NewDB()
	bookctrl := controller.NewBookController(db)
	app.GET("/books/", bookctrl.GetBooks)
	app.GET("/books/:id", bookctrl.GetBookByID)
	app.POST("/books/", bookctrl.CreateBook)
	app.PUT("/books/:id", bookctrl.UpdateBookByID)
	app.DELETE("/books/:id", bookctrl.DeleteBookByID)

	app.Run(":80")
}
