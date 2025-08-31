package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseService struct {
}

func (b *BaseService) success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func (b *BaseService) error(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, gin.H{
		"code": -1,
		"msg":  msg,
	})
}
