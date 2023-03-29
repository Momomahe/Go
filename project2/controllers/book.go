package controllers

import (
	"net/http"
	"project2/models"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) CreateBook(ctx *gin.Context) {
	book := models.Book{}
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.Name_book == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "name cannot be empty"})
		return
	}
	if err := c.masterDB.Create(&book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, book) //utk status code http 201
	//ctx.JSON(http.StatusOK, book)  //utk status code http 200
}
func (c *Controllers) GetBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	if bookID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid book ID"})
		return
	}

	book := models.Book{}
	err := c.masterDB.First(&book, bookID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
func (c *Controllers) GetAllBooks(ctx *gin.Context) {
	books := []models.Book{}
	err := c.masterDB.Find(&books).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(books) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no books found"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}
func (c *Controllers) UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	if bookID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid book ID"})
		return
	}

	book := models.Book{}
	err := c.masterDB.First(&book, bookID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.Name_book == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "name cannot be empty"})
		return
	}

	err = c.masterDB.Save(&book).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
func (c *Controllers) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	if bookID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid book ID"})
		return
	}

	book := models.Book{}
	err := c.masterDB.First(&book, bookID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	err = c.masterDB.Delete(&book).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted succesfully"})
}
