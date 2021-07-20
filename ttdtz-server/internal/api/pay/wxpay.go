package pay

import (
	"log"
	"net/http"
	"ttdtz-server/global"
	"ttdtz-server/internal/managers"
	"ttdtz-server/pkg/app"
	"ttdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func WxPay(c *gin.Context) {
	req := managers.WxPayRequestInfo{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	log.Printf("WxPay req => %+v", req)

	c.String(http.StatusOK, "")
	respdata, err := managers.WxPay(c.Request.Context(), &req)
	if err != nil {
		global.Logger.Error(err)
		response.ToErrorResponse(errcode.InvalidParams)
		//return
	}
	response.ToResponse(respdata)
}
