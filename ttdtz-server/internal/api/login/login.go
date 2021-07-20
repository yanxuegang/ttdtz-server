package login

import (
	"log"
	"net/http"
	"ttdtz-server/global"
	"ttdtz-server/internal/managers"

	//"ttdtz-server/internal/routers"
	"ttdtz-server/pkg/app"
	"ttdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type err error

// type params struct {
// 	OpenId   string `json:"open_id" form:"open_id"`
// 	Type     string `json:"type" form:"type"`
// 	Password string `json:"password" form:"password"`
// 	Channel  string `json:"channel" form:"channel"`
// }

// type loginRequestInfo struct {
// 	Cmd    int    `json:"cmd" form:"cmd"`
// 	Params params `json:"params" form:"params"`
// }

// type loginResponseInfo struct {
// 	Code         int    `json:"code"`
// 	OpenId       string `json:"open_id"`
// 	Sign         int    `json:"sign"`
// 	Money        int    `json:"money"`
// 	AccessToken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token"`
// }

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
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	//c.GetPostForm("ss")
	//formData := make(map[string]interface{})
	//respd := loginRequestInfo{}
	// 调用json包的解析，解析请求body
	//json.NewDecoder(c.Request.Body).Decode(&req)
	//for key, value := range respd {

	log.Printf("Login req => %+v", req)
	//}req

	c.String(http.StatusOK, "")

	//global.Logger.Infof("login req = %+v", req)
	//todo log
	//svc := managers.New(c.Request.Context())
	respdata, err := managers.Login(c.Request.Context(), &req)
	if err != nil {
		global.Logger.Error(err)
		response.ToErrorResponse(errcode.InvalidParams)
		//return
	}
	response.ToResponse(respdata)
	//return
}
