package account

import "time"

type Account struct {
	ID int32 
	Owner string
	Balance int64
	Currency string
	CreatedAt time.Time
	UpdatedAt time.Time 
}

func (a *Account) TableName() string {
	return "accounts"
  }