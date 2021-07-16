package global

import (
	"ttdtz-server/pkg/logger"
	"ttdtz-server/pkg/setting"
)

var (
	GlobalConfig    *setting.GlobalConfig
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
