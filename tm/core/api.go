package core

import (
	"tm/core/xdb"
	"tm/core/xfile"

	"github.com/gin-gonic/gin"
)

func InitApiv1(apiv1 *gin.RouterGroup) *gin.RouterGroup {
	{
		apiv1.POST("/db/desen", xdb.Find)
		apiv1.POST("/file/desen", xfile.Find)

	}
	return apiv1
}
