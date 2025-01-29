package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`

}

func GetBooks(c *gin.Context) {	
	var books map[string]Book

	if  err := client.NewRef("books/").Get(context.Background(), &books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	ref := client.NewRef("books/" + id)
	if err := ref.Get(context.Background(), &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ref := client.NewRef("books/")	
	if _, err := ref.Push(context.Background(), book); err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": ref.Key})
}

func UpdateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	ref := client.NewRef("books/" + id)
	if err := ref.Set(context.Background(), book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	ref := client.NewRef("books/" + id)
	if err := ref.Delete(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}