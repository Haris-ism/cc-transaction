package postgre

import (
	dbModels "cc-transaction/databases/postgresql/models"
	hModels "cc-transaction/hosts/callback/models"
	"cc-transaction/models"
	"cc-transaction/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	postgreDB struct {
		postgre *gorm.DB
	}
	PostgreInterface interface {
		Insert(req models.ItemList) error
		GetCC(req hModels.TransactionItems) (dbModels.CreditCards,error)
		OrderTransItem(req dbModels.Order) (dbModels.Order,error)
		DeductCC(req dbModels.CreditCards) error
		UpdateTransItem(req dbModels.Order) error
	}
)

func InitPostgre() PostgreInterface {
	host := utils.GetEnv("POSTGRE")
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Errorf("Failed to Init Postgre, Err:", err)
	} else {
		logrus.Printf("Init Postgre Success")
	}
	db.AutoMigrate(&models.ItemList{},&dbModels.Order{})

	return &postgreDB{
		postgre: db,
	}
}
