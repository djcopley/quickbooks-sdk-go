package model

import "time"

type Customer struct {
	CreditCardPaymentTxn struct {
		SyncToken            string `json:"SyncToken"`
		Domain               string `json:"domain"`
		CreditCardAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CreditCardAccountRef"`
		TxnDate     string `json:"TxnDate"`
		CurrencyRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CurrencyRef"`
		Amount         float64 `json:"Amount"`
		Sparse         bool    `json:"sparse"`
		BankAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"BankAccountRef"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"CreditCardPaymentTxn"`
	Time time.Time `json:"time"`
}
