package main

import (
	"example/web-service-gin/database"
	"example/web-service-gin/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.GetDb()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/module", router.GetModules(db))

	r.GET("/thread/:id", router.GetThreads(db))

	r.GET("/subcsribes/:moduleid", router.GetSubcsribers(db))
	r.POST("/subcsribes/:moduleid/:userid", router.Subcsribe(db))
	r.DELETE("/subcsribes/:moduleid/:userid", router.Unsubcsribe(db))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
