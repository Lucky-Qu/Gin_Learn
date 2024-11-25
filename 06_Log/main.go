package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	var err error
	g := gin.Default()
	//禁用颜色
	gin.DisableConsoleColor()
	//记录到文件
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	g.POST("/LogTest", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "执行成功",
		})
	})
	err = g.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
