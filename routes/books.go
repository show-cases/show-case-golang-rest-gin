package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiafangtao/showcases/model"
)

func addBooksRoute(rg *gin.RouterGroup) {
	// get all books
	var books []model.Book
	ping := rg.Group("/books")
	ping.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	})
}
