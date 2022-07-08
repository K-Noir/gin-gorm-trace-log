package controller

import (
	"hello/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func List(c *gin.Context) {
	ctx := c.Request.Context()
	logrus.WithContext(ctx).Info("Hello Web")
	service.List(ctx)
	c.JSON(http.StatusOK, time.Now().GoString())
}
