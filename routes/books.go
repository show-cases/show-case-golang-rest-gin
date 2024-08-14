package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jiafangtao/showcases/dal"
	"github.com/jiafangtao/showcases/model"
)

func addBooksRoute(rg *gin.RouterGroup) {
	// get all books
	var books []model.Book
	route := rg.Group("/books")
	route.GET("/", func(c *gin.Context) {

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

	route.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookId, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		//TODO: is the connection pooled?
		err = dal.Connect()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var book *model.Book
		book, err = dal.QueryBookById(bookId)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, book)
	})
}
