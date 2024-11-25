package main

import (
	"fmt"
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
		err = c.SaveUploadedFile(file, "./04_File"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		//c.JSON(200, gin.H{
		//	"msg": file.Filename,
		//})
		//将存储的图片返回到前端
		c.File("./04_File" + file.Filename)
	})
	err = g.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
