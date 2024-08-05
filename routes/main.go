package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	getRoutes()
	router.Run(":6060")
}

func getRoutes() {
	rg := router.Group("/")

	addPingRoute(rg)
	addBooksRoute(rg)
	addCommentsRoute(rg)
}
