package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiafangtao/showcases/model"
)

func addCommentsRoute(rg *gin.RouterGroup) {
	// get all books
	var comments []model.Comment
	ping := rg.Group("/comments")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
	})
}
