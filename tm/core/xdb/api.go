package xdb

import (
	"net/http"
	"tm/models"
	"tm/pkg/e"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	var config Config
	code := e.INVALID_PARAMS
	if c.BindJSON(&config) == nil {
		//将数据加入到异步队列
		if content, ok := config.str(); ok {
			models.Push("1", content)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})

	c.Abort()
}
