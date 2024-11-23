package main

import "github.com/gin-gonic/gin"

func main() {
	var err error
	g := gin.Default()
	BindWithCheck(g)
	//BindWithoutCheck(g)
	err = g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
