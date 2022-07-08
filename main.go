package main

import (
	"hello/controller"
	"hello/logger"
	"hello/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 打印日志时，取ctx中的trace_id
	logrus.AddHook(utils.NewLogTrace())

	router := gin.Default()

	// http请求时，为ctx添加trace_id
	router.Use(logger.WithTrace())

	router.GET("/book", controller.List)
	err := router.Run(":8070")
	if err != nil {
		panic(err)
	}
}
