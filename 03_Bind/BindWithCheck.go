package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func BindWithCheck(g *gin.Engine) {
	g.POST("/BindWithCheck", func(c *gin.Context) {
		var err error
		var U = UserWithCheck{}
		if Age, ok := binding.Validator.Engine().(*validator.Validate); ok {
			err = Age.RegisterValidation("AgeCheck", AgeCheck)
			if err != nil {
				panic(err)
			}
		}
		err = c.ShouldBindJSON(&U)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"Data": U,
		})
	})
}
func AgeCheck(fl validator.FieldLevel) bool {
	if age, ok := fl.Field().Interface().(int); ok {
		if age < 18 {
			return false
		}
	}
	return true
}
