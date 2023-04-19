package model

type BankAcc struct {
	AccountId         uint   `json:"account_id"`
	UserId            int    `json:"user_id"`
	BankName          string `json:"bank_name"`
	AccountNumber     string `json:"account_number"`
	AccountHolderName string `json:"account_holder_name"`
}