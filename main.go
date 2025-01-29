package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitFirebase()
	r := gin.Default()
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBookByID)
	r.POST("/books", CreateBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
	r.Run(":8080")
}