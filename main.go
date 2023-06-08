package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func setRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "first",
			"q":       c.Query("q"),
		})
	})

	router.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("name: %s; message: %s", name, message)
		c.JSON(202, gin.H{
			"name":    name,
			"message": message,
		})
	})
}

func main() {
	router := gin.Default()
	setRouter(router)
	_ = router.Run()
}
