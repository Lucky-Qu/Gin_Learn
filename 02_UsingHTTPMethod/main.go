package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.Query("user")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"id":    id,
			"user":  user,
			"pwd":   pwd,
			"Reply": "GET测试成功！",
		})
	})
	r.POST("/test", func(c *gin.Context) {
		user := c.PostForm("user")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user":  user,
			"pwd":   pwd,
			"Reply": "POST测试成功！",
		})
	})
	r.DELETE("/test/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id":    id,
			"Reply": "DELETE测试成功！",
		})
	})
	r.PUT("/test", func(c *gin.Context) {
		user := c.PostForm("user")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user":  user,
			"pwd":   pwd,
			"Reply": "PUT测试成功！",
		})
	})
	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
