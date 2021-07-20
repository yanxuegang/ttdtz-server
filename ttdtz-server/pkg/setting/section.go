package setting

import (
	"time"
)

type GlobalConfig struct {
	Server   ServerSettingS
	App      AppSettingS
	Database DatabaseSettingS
	Redis    RedisSettingS
	Wx       WxSettingS
}

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type WxSettingS struct {
	AppId                string
	AppSecret            string
	MchId                string
	WxPayChechUrl        string
	ApiPrefix            string
	AuthUrlFormat        string
	AccessTokenUrlFormat string
}

type DBConn struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type DatabaseSettingS struct {
	DBConns []DBConn
}

type RedisConn struct {
	Name               string
	Host               string
	Port               string
	Password           string
	PoolSize           int
	PoolTimeout        int
	IdleTimeout        int
	IdleCheckFrequency uint32
	Database           int
}

type RedisCache struct {
	Conns      []RedisConn
	Expiration int
}

type RedisSettingS struct {
	Cache RedisCache
}

func (s *Setting) ReadSection(v interface{}) error {
	//todu 各个config分割下
	err := s.vp.Unmarshal(v)
	if err != nil {
		return err
	}
	return nil
	// err := s.vp.UnmarshalKey(k, v)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
