package usecase

import (
	"cc-transaction/models"
)

func (uc *usecase) WriteRedis(req models.RedisReq) error {
	err := uc.redis.WriteRedis(req)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) ReadRedis(req models.RedisReq) (string, error) {
	res, err := uc.redis.ReadRedis(req)
	if err != nil {
		return res, err
	}

	return res, nil
}
