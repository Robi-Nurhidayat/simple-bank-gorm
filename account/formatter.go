package account

import "time"

type AcountFormatter struct {
	ID        int32     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AccountFormatter(account Account) AcountFormatter {

	var formatter AcountFormatter
	formatter.ID = account.ID
	formatter.Owner = account.Owner
	formatter.Balance = account.Balance
	formatter.Currency = account.Currency
	formatter.CreatedAt = account.CreatedAt
	formatter.UpdatedAt = account.UpdatedAt

	return formatter
}

func ListAccountFormatter(accounts []Account) []AcountFormatter {

	var listAccounts []AcountFormatter

	for _, v := range accounts {
		listAccounts = append(listAccounts, AccountFormatter(v))
	}

	return listAccounts
}
