package tool

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

/*
	RedisStore结构体实现了Store接口的所有方法
*/
var Rds RedisStore

func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig
	// 创建Redis的连接
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})

	Rds = RedisStore{
		client: client,
	}
	return &Rds
}

func (rs *RedisStore) Set(id string, value string) {
	if err := rs.client.Set(id, value, time.Minute*10).Err(); err != nil {
		log.Println(err)
	}
}

func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	return false
}
