package controller

import (
	"database/sql"
	"net/http"

	"testing/model"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	db *sql.DB
}

func NewBookController(db *sql.DB) BookController {
	return BookController{
		db: db,
	}
}

func (ctrl BookController) GetBooks(c *gin.Context) {
	books := []model.Book{}

	rows, err := ctrl.db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		book := model.Book{}

		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Descr)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
	}
	
	c.JSON(http.StatusCreated, books) //201
	//c.JSON(http.StatusOK, books) //200
}
func (ctrl BookController) CreateBook(c *gin.Context) {
	book := model.Book{}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := ctrl.db.Exec("INSERT INTO books (title, author, descr) VALUES ($1, $2, $3)", book.Title, book.Author, book.Descr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}
func (ctrl BookController) GetBookByID(c *gin.Context) {
	bookID := c.Param("id")

	book := model.Book{}

	err := ctrl.db.QueryRow("SELECT * FROM books WHERE id=$1", bookID).Scan(&book.BookID, &book.Title, &book.Author, &book.Descr)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}
func (ctrl BookController) UpdateBookByID(c *gin.Context) {
	bookID := c.Param("id")

	book := model.Book{}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.db.Exec("UPDATE books SET title=$1, author=$2, descr=$3 WHERE id=$4", book.Title, book.Author, book.Descr, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
func (ctrl BookController) DeleteBookByID(c *gin.Context) {
	bookID := c.Param("id")

	result, err := ctrl.db.Exec("DELETE FROM books WHERE id=$1", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
