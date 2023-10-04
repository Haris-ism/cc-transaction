package postgre

import (
	"cc-transaction/models"

	"github.com/sirupsen/logrus"
)

func (db *postgreDB) Insert(req models.ItemList) error {
	if err := db.postgre.Create(&req).Error; err != nil {
		logrus.Errorf("Failed to Insert Postgre, Err:", err)
		return err
	}
	return nil
}
