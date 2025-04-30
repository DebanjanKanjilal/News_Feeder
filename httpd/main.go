package main

import (
	"news-feeder/httpd/handler"

	"github.com/gin-gonic/gin"
)

// main function
func main() {
	// creates new gin router instance and attach default middleware and logger
	r := gin.Default()

	// defines the gateway "/ping" and prepares a response in form of "message" : "pong"
	r.GET("/ping", handler.PingGet)

	// runs the server
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
