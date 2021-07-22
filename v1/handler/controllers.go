package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sandiptorane/gin_poc/model"
	"log"
	"net/http"
	"strconv"
)

type BookStore struct {
	books []model.Book
}

func (b *BookStore) CreateBook(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(b.books) != 0 {
		book.Id = b.books[len(b.books)-1].Id + 1
	}
	b.books = append(b.books, book)
	log.Println("user_id:", userId, "new book is created")
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (b *BookStore) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	name := c.Query("name")
	for _, book := range b.books {
		if name != "" {
			if book.Id == id && book.Name == name {
				c.JSON(http.StatusOK, gin.H{"data": book})
				return
			}
		} else if book.Id == id {
			c.JSON(http.StatusOK, gin.H{"data": book})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "requested book not found"})
}

func (b *BookStore) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": b.books})
}

func (b *BookStore) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var bookForUpdate model.Book
	if err := c.BindJSON(&bookForUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, book := range b.books {
		if book.Id == id {
			b.books[i].Author = bookForUpdate.Author
			b.books[i].Name = bookForUpdate.Name
			b.books[i].Description = bookForUpdate.Description
			c.JSON(http.StatusOK, "book updated successfully")
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "requested book not found"})
}

func (b *BookStore) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	for i, book := range b.books {
		if book.Id == id {
			//if book id found delete it
			b.books = append(b.books[:i], b.books[i+1:]...)
			c.JSON(http.StatusOK, "book deleted successfully")
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "requested book not found"})
}
