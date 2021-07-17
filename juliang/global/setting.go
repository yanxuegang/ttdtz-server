package global

import (
	"juliang/pkg/logger"
	"juliang/pkg/setting"
)

var (
	GlobalConfig    *setting.GlobalConfig
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
