package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "hello wts item testando"})
    })

	r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "bora carai"})
    })

    r.Run(":8000")
}