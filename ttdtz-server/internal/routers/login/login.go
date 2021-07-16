package login

import (
	"ttdtz-server/global"
	"ttdtz-server/internal/managers"
	"ttdtz-server/pkg/app"
	"ttdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type err error

type params struct {
	OpenId   string `json:"open_id"`
	Type     string `json:"type"`
	Password string `json:"password"`
	Channel  string `json:"channel"`
}

type loginRequestInfo struct {
	Cmd    int    `json:"cmd"`
	Params params `json:"params"`
}

type loginResponseInfo struct {
	OpenId string `json:"open_id"`
	Sign   int    `json:"sign"`
	Money  int    `json:"money"`
}

// @Summary 登录
// @Produce json
// @Param loginRequestInfo query loginRequestInfo true "loginRequestInfo"
// @Param Params query params true "params"
// @Success 200 {object} loginResponseInfo "成功"
// @Failure 400 {object} err "请求错误"
// @failure 500 {object} err "内部错误"
// @Router /api/v1/login [get]
func Login(c *gin.Context) {
	req := managers.LoginRequest{}
	response := app.NewResponse(c)
	c.ShouldBind(&req)
	global.Logger.Infof("login req = %+v", req)
	//todo log
	//svc := managers.New(c.Request.Context())
	respdata, err := managers.Login(c.Request.Context(), &req)
	if err != nil {
		global.Logger.Error(err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	response.ToResponse(respdata)
	return
}
