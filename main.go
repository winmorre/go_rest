package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry porter", Author: "J winmorre"},
	{ID: "2", Title: "Rich Dad Poor Dad", Author: "MOn D."},
	{ID: "3", Title: "Babylons Kingpins", Author: "J.F frost"},
}

func main() {

	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world web",
		})
	})

	r.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, books)
	})
	r.Run()

	r.POST("/book", func(ctx *gin.Context) {
		var book Book
		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		books = append(books, book)
		ctx.JSON(http.StatusCreated, books)
	})
}
