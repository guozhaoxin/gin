package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main () {
	r := gin.Default()
	r.POST("/users", func (context *gin.Context) {
		context.JSON(http.StatusOK, "users")
	})
	/* ... ... */
	r.Run(":8080")
}