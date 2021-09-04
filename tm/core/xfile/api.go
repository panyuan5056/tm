package xfile

import (
	"tm/models"
	"tm/pkg/e"
	"tm/pkg/logging"

	"net/http"
	"tm/pkg/util"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	var form UploadForm
	code := e.INVALID_PARAMS
	if err := c.ShouldBind(&form); err == nil {
		code = e.SUCCESS
		if ext, ok := form.valid(); ok {
			dst := util.ParseFile(ext)
			if err := c.SaveUploadedFile(form.Upload, dst); err != nil {
				logging.Error(err.Error())
			} else {
				if content, ok := form.str(dst); ok {
					models.Push("2", content)
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})
	c.Abort()
}
