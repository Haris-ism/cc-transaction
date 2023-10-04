package redis_db

import (
	"context"
	"cc-transaction/models"
	"time"

	"github.com/sirupsen/logrus"
)

func (uc *redisDB) WriteRedis(req models.RedisReq) error {
	// store data using SET command
	err := uc.redis.Set(context.Background(), req.Key, req.Data, time.Duration(req.Exp)*time.Second).Err()
	if err != nil {
		logrus.Errorf("Error on :%v, unable to SET data. error: %v", "WriteRedis DB Function", err)
		return err
	}
	logrus.Printf("set operation success on WriteRedis DB Function")

	return nil
}

func (uc *redisDB) ReadRedis(req models.RedisReq) (string, error) {

	// get data
	res, err := uc.redis.Get(context.Background(), req.Key).Result()
	if err != nil {
		logrus.Errorf("Error on :%v, unable to GET data. error: %v", "ReadRedis DB Function", err)
		return res, err
	}
	logrus.Printf("get operation success on ReadRedis DB Function. result:", res)

	return res, nil
}
