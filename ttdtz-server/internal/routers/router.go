package routers

import (
	_ "ttdtz-server/docs"
	"ttdtz-server/internal/routers/login"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/login", login.Login)
		apiv1.POST("/wxlogin", login.WxLogin)
	}
	return r
}
