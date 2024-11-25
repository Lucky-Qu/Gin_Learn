package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middle1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "中间件一号执行前",
		})
		c.Next()
		c.JSON(200, gin.H{
			"msg": "中间件一号执行后",
		})
	}
}
func middle2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "中间件二号执行前",
		})
		c.Next()
		c.JSON(200, gin.H{
			"msg": "中间件二号执行后",
		})
	}
}
func main() {
	g := gin.Default()
	v1 := g.Group("/v1")
	v1.Use(middle1()).Use(middle2())
	v1.POST("MiddlewareAndGroupTest", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "中间件和组的测试",
		})
	})
	err := g.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
