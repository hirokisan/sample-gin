package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name/", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": name,
		})
	})
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "guest")
		lastname := c.Query("lastname")
		c.JSON(200, gin.H{
			"message":   "OK",
			"firstname": firstname,
			"lastname":  lastname,
		})
	})
	r.POST("/user/", func(c *gin.Context) {
		name := c.PostForm("name")
		c.JSON(200, gin.H{
			"message": "OK",
			"name":    name,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
