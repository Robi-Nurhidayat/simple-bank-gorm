package account

import "gorm.io/gorm"

type AccountRepository interface {
	CreateAccount(account Account) (Account, error)
	ListAccounts()([]Account,error)
	UpdateAccount(account Account)(Account,error)
	GetAccountByID(id int32)(Account,error)
	DeleteById(id int32)(error)
}

type repository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &repository{db: db}
}

func (r *repository) CreateAccount(account Account) (Account, error) {

	err := r.db.Create(&account).Error

	if err != nil {
		return account, err
	}

	return account, nil

}


func (r *repository) ListAccounts() ([]Account, error) {

	var accounts []Account

	err := r.db.Find(&accounts).Error

	if err != nil {
		return accounts,err
	}

	return accounts,nil

	
}

func (r *repository) UpdateAccount(account Account) (Account, error) { 

	err := r.db.Save(&account).Error

	if err != nil {
		return account,err
	}

	return account,nil

	
}

func (r *repository) GetAccountByID(id int32) (Account, error) {

	var account Account

	err := r.db.Where("id = ?", id).First(&account).Error

	if err != nil {
		return account,err
	}

	return account,nil

 }


 func (r *repository) DeleteById(id int32)  error {

	err := r.db.Delete(&Account{}," id = ? ", id).Error

	if err != nil {
		return err
	}

	return nil
 }