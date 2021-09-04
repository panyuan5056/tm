package auth

import (
	"net/http"

	"tm/models"
	"tm/pkg/e"

	"tm/pkg/logging"
	"tm/pkg/util"

	"github.com/gin-gonic/gin"
)

func Online(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": e.SUCCESS,
		"msg":    e.GetMsg(e.SUCCESS),
	})
	c.Abort()
}

func Register(c *gin.Context) {
	var auth models.Auth
	code := e.INVALID_PARAMS
	if c.BindJSON(&auth) == nil {
		newPassword := models.GenPassword(auth.Password)
		auth.Password = newPassword
		auth.Active = 1
		models.DB.Create(&auth)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func Auth(c *gin.Context) {
	var auth models.Auth
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if c.BindJSON(&auth) == nil {
		isExist := models.CheckAuth(auth.Username, auth.Password)
		if isExist {
			token, err := util.GenerateToken(auth.Username, auth.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
			logging.Error(e.GetMsg(code))
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
