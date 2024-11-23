package main

import "github.com/gin-gonic/gin"

func BindWithoutCheck(g *gin.Engine) {
	g.POST("/BindWithoutCheck", func(c *gin.Context) {
		var u UserWithoutCheck
		var err error
		err = c.ShouldBindJSON(&u)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"data": u,
		})
	})
}
