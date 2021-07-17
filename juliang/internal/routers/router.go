package routers

import (
	_ "juliang/docs"
	"juliang/internal/routers/login"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("/login", login.Login)
		loginGroup.POST("/wxlogin", login.WxLogin)
	}
	payGroup := r.Group("/pay")
	{
		payGroup.POST("/login", login.Login)
		payGroup.POST("/wxlogin", login.WxLogin)
	}
	adGroup := r.Group("/ad")
	{
		adGroup.POST("/login", login.Login)
		adGroup.POST("/wxlogin", login.WxLogin)
	}
	return r
}
