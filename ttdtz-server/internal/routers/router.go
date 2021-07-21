package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	_ "ttdtz-server/docs"
	"ttdtz-server/internal/api"
	"ttdtz-server/internal/api/login"
	"ttdtz-server/internal/api/pay"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type RequestInfo struct {
	Cmd    int         `json:"cmd"`
	Params interface{} `json:"params"`
}

var (
	process *api.Processor
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api")
	{
		apiv1.POST("/v1", Router)
		apiv1.OPTIONS("/v1", Router)
		apiv1.GET("/v1", Router)

		apiv1.POST("/login", login.Login)
		apiv1.OPTIONS("/login", login.Login)
		apiv1.GET("/login", login.Login)

		apiv1.POST("/wxlogin", login.WxLogin)
		apiv1.OPTIONS("/wxlogin", login.WxLogin)
		apiv1.GET("/wxlogin", login.WxLogin)

		apiv1.POST("/wxpay", pay.WxPay)
		apiv1.OPTIONS("/wxpay", pay.WxPay)
		apiv1.GET("/wxpay", pay.WxPay)
	}
	return r
}

func Router(c *gin.Context) {
	context.WithValue(c.Request.Context(), "test", "yxg")
	//formData := make(map[string]interface{})
	//var bodyBytes []byte
	respd := RequestInfo{}
	//if c.Request.Body != nil {
	//bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	json.NewDecoder(c.Request.Body).Decode(&respd)
	//for key, value := range formData {
	log.Printf("Login req => %+v", respd)
	// c.Keys = make(map[string]interface{})
	// c.Keys["cmd"] = respd.Cmd
	// c.Keys["params"] = respd.Params
	c.Request.Header.Set("cmd", "1001")
	c.Set("cmd", respd.Cmd)
	c.Set("params", respd.Params)
	//}
	// 把刚刚读出来的再写进去
	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	api.GetProcessor().MessageInfo[api.MessageID(respd.Cmd)].Handler(c)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	c.JSON(http.StatusOK, "Options Request!")
		//}

		c.Next()
	}
}
