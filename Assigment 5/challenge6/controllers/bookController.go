package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.BookID = fmt.Sprintf("%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.String(http.StatusCreated, "\"Created\"")
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "Data Not Found",
			"error messagae": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.String(http.StatusOK, "\"Updated\"")
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "Data Not Found",
			"error messagae": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "Data Not Found",
			"error messagae": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.String(http.StatusOK, "\"Deleted\"")
}
func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": BookDatas,
	})
}
