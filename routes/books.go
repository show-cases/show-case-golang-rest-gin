package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiafangtao/showcases/dal"
	"github.com/jiafangtao/showcases/model"
)

func addBooksRoute(rg *gin.RouterGroup) {
	// get all books
	var books []model.Book
	ping := rg.Group("/books")
	ping.GET("/", func(c *gin.Context) {

		err := dal.Connect()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		books, err = dal.QueryAllBooks()
		fmt.Println("books: ", len(books))
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	})
}
