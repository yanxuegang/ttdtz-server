package app

import (
	"fmt"
	"log"
	"ttdtz-server/global"
)

/********************* sdk相关定义 *********************/
func GetWxAuthUrl(code string) string {
	log.Printf("GetWxAuthUrl = %s", fmt.Sprintf(global.GlobalConfig.Wx.AuthUrlFormat,
		global.GlobalConfig.Wx.ApiPrefix,
		global.GlobalConfig.Wx.AppId,
		global.GlobalConfig.Wx.AppSecret,
		code))
	return fmt.Sprintf(global.GlobalConfig.Wx.AuthUrlFormat,
		global.GlobalConfig.Wx.ApiPrefix,
		global.GlobalConfig.Wx.AppId,
		global.GlobalConfig.Wx.AppSecret,
		code)
}
