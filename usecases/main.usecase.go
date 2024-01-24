package usecase

import (
	con "cc-transaction/controllers/models"
	host "cc-transaction/hosts"
	cModels "cc-transaction/hosts/callback/models"
	hm "cc-transaction/hosts/merchant/models"
	"cc-transaction/models"

	postgre "cc-transaction/databases/postgresql"
	redis_db "cc-transaction/databases/redis"
)

type (
	usecase struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
		host	host.HostInterface
	}
	UsecaseInterface interface {
		WriteRedis(models.RedisReq) error
		ReadRedis(req models.RedisReq) (string, error)
		InsertDB(req models.ItemList) error
		InquiryItems()([]hm.InquiryItems,error)
		InquiryDiscounts()([]hm.InquiryDiscounts,error)
		TransItem(req cModels.DecTransactionItems, headers con.ReqHeader)(string,error)
	}
)

func InitUsecase(postgre postgre.PostgreInterface, redis redis_db.RedisInterface, host host.HostInterface) UsecaseInterface {
	return &usecase{
		postgre: postgre,
		redis:   redis,
		host: host,
	}
}
