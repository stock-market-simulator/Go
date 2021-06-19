package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CICD branch test2",
		})
	})
	r.Run(":5000")
}
