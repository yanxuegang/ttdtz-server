package global

import (
	"juliang/internal/rmodels"
	"time"

	"github.com/go-redis/redis"
)

type CacheZ = redis.Z

var (
	CacheConnStrategy rmodels.CacheConnStrategy
	cacheLoading      bool
	cacheLoaded       bool
)

func CacheIncr(key string) (int64, error) {
	return CacheConnStrategy.GetClient("cache").Incr(key).Result()
}

func CacheSet(key string, value interface{}) (string, error) {
	return CacheConnStrategy.GetClient("cache").Set(key, value, CacheConnStrategy.GetCacheExpiration()).Result()
}

func CacheSetWithExpiration(key string, value interface{}, expiration time.Duration) (string, error) {
	return CacheConnStrategy.GetClient("cache").Set(key, value, expiration).Result()
}

func CacheGet(key string) (string, error) {
	return CacheConnStrategy.GetClient("cache").Get(key).Result()
}

func CacheExists(key string) (int64, error) {
	return CacheConnStrategy.GetClient("cache").Exists(key).Result()
}

func CacheDel(keys ...string) (int64, error) {
	return CacheConnStrategy.GetClient("cache").Del(keys...).Result()
}

func CacheSAdd(key string, members ...interface{}) (int64, error) {
	return CacheConnStrategy.GetClient("cache").SAdd(key, members...).Result()
}

func CacheSRem(key string, members ...interface{}) (int64, error) {
	return CacheConnStrategy.GetClient("cache").SRem(key, members...).Result()
}

func CacheSMembers(key string) ([]string, error) {
	return CacheConnStrategy.GetClient("cache").SMembers(key).Result()
}

func CacheSCard(key string) (int64, error) {
	return CacheConnStrategy.GetClient("cache").SCard(key).Result()
}

func CacheHMGet(key string, fields []string) ([]interface{}, error) {
	return CacheConnStrategy.GetClient("cache").HMGet(key, fields...).Result()
}

func CacheHDel(key string, fields []string) (int64, error) {
	return CacheConnStrategy.GetClient("cache").HDel(key, fields...).Result()
}

func CacheHMSet(key string, fields map[string]interface{}) (string, error) {
	return CacheConnStrategy.GetClient("cache").HMSet(key, fields).Result()
}

func CacheHGetAll(key string) (map[string]string, error) {
	return CacheConnStrategy.GetClient("cache").HGetAll(key).Result()
}

func CacheZAdd(key string, members ...CacheZ) (int64, error) {
	return CacheConnStrategy.GetClient("cache").ZAdd(key, members...).Result()
}

func CacheZScore(key string, member string) (float64, error) {
	return CacheConnStrategy.GetClient("cache").ZScore(key, member).Result()
}
