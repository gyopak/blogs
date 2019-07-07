package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./client", true)))


	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
