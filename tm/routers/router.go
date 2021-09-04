package routers

import (
	"tm/core"
	"tm/middleware/jwt"
	"tm/pkg/setting"
	"tm/routers/auth"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	r.POST("/online", auth.Online)
	r.POST("/auth", auth.Auth)
	r.POST("/register", auth.Register)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	apiv1 = core.InitApiv1(apiv1)

	return r
}
