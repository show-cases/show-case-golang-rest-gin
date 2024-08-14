package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiafangtao/showcases/dal"
	"github.com/jiafangtao/showcases/model"
)

type CommentsQuery struct {
	BookId int `form:"bookId"`
}

func addCommentsRoute(rg *gin.RouterGroup) {
	// get all books
	var comments []model.Comment
	router := rg.Group("/comments")
	router.GET("/", func(c *gin.Context) {

		var query CommentsQuery
		err := c.ShouldBind(&query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		bookId := query.BookId
		fmt.Println("bookId=", bookId)

		err = dal.Connect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if bookId != 0 {
			comments, err = dal.QueryCommentsByBookId(bookId)
		} else {
			comments, err = dal.QueryAllComments()
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(comments) == 0 {
			comments = []model.Comment{}
		}

		c.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
	})
}
