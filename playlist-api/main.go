package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/playlists", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run()
}

func getPlaylists(c *gin.Context) {

}
