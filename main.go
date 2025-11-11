package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/webservice", handlerRequest)
	router.GET("/webservice", getURLs)

	router.Run("localhost:1010")
}
