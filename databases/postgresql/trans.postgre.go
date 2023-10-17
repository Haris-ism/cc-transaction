package postgre

import (
	dbModels "cc-transaction/databases/postgresql/models"
	"cc-transaction/hosts/callback/models"
	"fmt"
)

func(db *postgreDB)GetCC(req models.TransactionItems) (dbModels.CreditCards,error){
	fmt.Println("masuk db teu?")
	reqDB:=dbModels.CreditCards{}
	err:=db.postgre.Where("cc_number = ?",req.CCNumber).Find(&reqDB).Error
	if err!=nil{
		return reqDB,err
	}
	return reqDB,nil
}

func(db *postgreDB)OrderTransItem(req dbModels.Order) (dbModels.Order,error){
	err:=db.postgre.Create(&req).Error
	if err!=nil{
		return req,err
	}
	return req,nil
}

func(db *postgreDB)DeductCC(req dbModels.CreditCards) error{
	err:=db.postgre.Save(&req).Error
	if err!=nil{
		return err
	}
	return nil
}
func(db *postgreDB)UpdateTransItem(req dbModels.Order) error{
	err:=db.postgre.Save(&req).Error
	if err!=nil{
		return err
	}
	return nil
}
