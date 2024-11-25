package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var err error
	g := gin.Default()
	g.POST("/FileTest", func(c *gin.Context) {
		//接收文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//存储文件到本地
		err = c.SaveUploadedFile(file, "./"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"msg": file.Filename,
		})
	})
	err = g.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
