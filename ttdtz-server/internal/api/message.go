package api

import (
	"fmt"
	"math"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	_                   = iota
	MessageRequestLogin = 1000 + iota
	MessageRequestWxLogin
)

/*
 * 通信处理器(单例)
 */
type Processor struct {
	MessageInfo map[MessageID]*MessageInfo
	MessageID   map[reflect.Type]MessageID
}

/*
 * 消息
 */
type MessageInfo struct {
	Type reflect.Type
	//Handler      handler.HandlerInterface
	Handler      func(*gin.Context)
	ResponseType reflect.Type
}
type MessageID uint16

/*
 * 消息处理器
 */
// type MessageHandler func([]interface{}) (interface{}, error)

func (p *Processor) Register(id MessageID, msg interface{}) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		fmt.Errorf("message pointer required")
	}
	if _, ok := p.MessageID[msgType]; ok {
		fmt.Errorf("Message %v is already registered", msgType)
	}
	if len(p.MessageInfo) >= math.MaxUint16 {
		fmt.Errorf("Too many messages (max = %v)", math.MaxUint16)
	}

	msgInfo := new(MessageInfo)
	msgInfo.Type = msgType

	p.MessageInfo[id] = msgInfo
	p.MessageID[msgType] = id
}

func (p *Processor) SetHandler(msg interface{}, handler func(*gin.Context)) {
	msgType := reflect.TypeOf(msg)
	id, ok := p.MessageID[msgType]
	if !ok {
		fmt.Errorf("Message %v not registered", msgType)
	}
	handlerType := reflect.TypeOf(handler)
	if handlerType == nil || handlerType.Kind() != reflect.Ptr {
		fmt.Errorf("Handler pointer required")
	}
	p.MessageInfo[id].Handler = handler
}
