package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func initLogrus() {
	path := "./06_Log/LogTest.log"
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.WithError(err).Fatal("打开日志文件失败")
	}

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(log.InfoLevel)
}

func middleLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		StartTime := time.Now()
		c.Next()
		EndTime := time.Now()
		initLogrus()
		log.WithFields(log.Fields{
			"StartTime": StartTime,
			"EndTime":   EndTime,
		}).Info("处理请求")
	}
}
func main() {
	var err error

	g := gin.Default()
	v1 := g.Group("/v1")
	v1.Use(middleLog())
	v1.POST("/LogTest", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "执行成功",
		})
	})
	err = g.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
