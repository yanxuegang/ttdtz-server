package rmodels

import (
	"fmt"
	"juliang/pkg/randutil"
	"juliang/pkg/setting"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const (
	CacheNil = redis.Nil
)

// cacheConn Cache连接
// @param PoolSize 连接池连接数, 0表示默认值(CPU数量*10)
// @param PoolTimeout 连接池申请连接的超时时间,单位为秒,0表示默认值(ReadTimeout+1)
// @param IdleTimeout 空闲连接超时时间,单位为秒,0表示默认值(5分钟)
// @param IdleCheckFrequency 空闲连接检查频率,单位为秒,0表示默认值(1分钟)
type cacheConn struct {
	Name               string
	Host               string
	Port               string
	Password           string
	PoolSize           int           `json:"pool_size"`
	PoolTimeout        time.Duration `json:"pool_timeout"`
	IdleTimeout        time.Duration `json:"idle_timeout"`
	IdleCheckFrequency time.Duration `json:"idle_check_frequency"`
	Database           int

	Client *redis.Client
	Ring   *redis.Ring
}

func (m *cacheConn) Init() {
	opt := &redis.Options{
		Addr:               m.Host,
		Password:           m.Password,
		DB:                 m.Database,
		PoolSize:           m.PoolSize,
		PoolTimeout:        time.Second * m.PoolTimeout,
		IdleTimeout:        time.Second * m.IdleTimeout,
		IdleCheckFrequency: time.Second * m.IdleCheckFrequency,
	}
	m.Client = redis.NewClient(opt)
	if _, err := m.Client.Ping().Result(); err != nil {
		panic(fmt.Sprintf("[Cache][Exception] %s cache connect failed: %s ...", m.Name, err))
	}
	log.Printf("[Cache][Success] %s cache connected ...", m.Name)

	//if global.GlobalConfig.Server.RunMode == "debug" {
	m.Client.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			argsStr := ""
			for _, v := range cmd.Args() {
				switch v.(type) {
				case string:
					argsStr += fmt.Sprintf("%s ", v)
				case []byte:
					argsStr += fmt.Sprintf("%s ", string(v.([]byte)))
				case int64:
					argsStr += fmt.Sprintf("%d ", v.(int64))
				}
			}
			log.Printf("[Cache]Processing: <%s>", argsStr)
			err := old(cmd)
			//fmt.Printf("[Cache]finished processing: <%s>\n", cmd)
			return err
		}
	})
	//}
}

// Cache
// @param Expiration 过期时间,单位为秒,0代表不过期
type cache struct {
	Conns      []*cacheConn
	Expiration time.Duration
}

// cacheConnStrategy Cache连接策略
type CacheConnStrategy struct {
	Cache cache
}

func (m CacheConnStrategy) Init() {
	for _, conn := range m.Cache.Conns {
		conn.Init()
	}
}

func (m CacheConnStrategy) GetClient(name string) *redis.Client {
	var conns []*cacheConn

	conns = m.Cache.Conns

	// 因为使用Convus，所以具体在那个Redis实例中不需要代码控制
	if len(conns) > 0 {
		idx, _ := randutil.IntRange(0, len(conns)-1)
		return conns[idx].Client
	}
	return nil
}

func (m CacheConnStrategy) GetCacheExpiration() time.Duration {
	return m.Cache.Expiration * time.Second
}

type CacheZ = redis.Z

var (
	defaultCacheConnStrategy CacheConnStrategy
	cacheLoading             bool
	cacheLoaded              bool
)

func NewCache(setting *setting.RedisSettingS) CacheConnStrategy {
	InitCache(setting)
	return defaultCacheConnStrategy
}

func InitCache(setting *setting.RedisSettingS) {
	if cacheLoading || cacheLoaded {
		return
	}
	cacheLoading = true
	defer func() {
		cacheLoading = false
	}()
	for _, conn := range setting.Cache.Conns {
		v := &cacheConn{
			Name:               conn.Name,
			Host:               conn.Host,
			Port:               conn.Port,
			Password:           conn.Password,
			PoolSize:           conn.PoolSize,
			PoolTimeout:        time.Duration(conn.PoolTimeout),
			IdleTimeout:        time.Duration(conn.IdleTimeout),
			IdleCheckFrequency: time.Duration(conn.IdleCheckFrequency),
			Database:           conn.Database,
		}
		defaultCacheConnStrategy.Cache.Conns = append(defaultCacheConnStrategy.Cache.Conns, v)
	}
	defaultCacheConnStrategy.Cache.Expiration = time.Duration(setting.Cache.Expiration)
	defaultCacheConnStrategy.Init()
	cacheLoaded = true

	if !cacheLoaded {
		panic(fmt.Sprint("[Cache][Exception] redis init error ..."))
	}
}
