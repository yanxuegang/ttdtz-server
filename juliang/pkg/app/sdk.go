package app

import (
	"fmt"
	"juliang/global"
)

/********************* sdk相关定义 *********************/
func GetWxAuthUrl(code string) string {
	return fmt.Sprintf(global.GlobalConfig.Wx.AuthUrlFormat,
		global.GlobalConfig.Wx.ApiPrefix,
		global.GlobalConfig.Wx.AppId,
		global.GlobalConfig.Wx.AppSecret,
		code)
}
