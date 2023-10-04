package redis_db

import (
	"cc-transaction/models"
	"cc-transaction/utils"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type (
	redisDB struct {
		redis *redis.Client
	}
	RedisInterface interface {
		WriteRedis(models.RedisReq) error
		ReadRedis(req models.RedisReq) (string, error)
	}
)

func InitRedis() RedisInterface {
	host := utils.GetEnv("REDIS_HOST")
	password := utils.GetEnv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	logrus.Printf("Init Redis Success")

	return &redisDB{
		redis: client,
	}
}
