package api

import (
	"reflect"
	"ttdtz-server/internal/api/login"
	"ttdtz-server/internal/managers"
)

// Processor 消息处理器
var (
	processor Processor
)

func GetProcessor() *Processor {
	return &processor
}

// MessageRegister 消息注册
func MessageRegister() {
	//processor := new(Processor)
	processor.MessageID = make(map[reflect.Type]MessageID)
	processor.MessageInfo = make(map[MessageID]*MessageInfo)

	processor.Register(MessageRequestLogin, managers.LoginRequest{})
	processor.SetHandler(managers.LoginRequest{}, login.WxLogin)

}
