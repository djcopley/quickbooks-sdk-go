package model

import "time"

type TaxPayment struct {
	TaxPayment struct {
		Refund            string `json:"Refund"`
		SyncToken         string `json:"SyncToken"`
		Domain            string `json:"domain"`
		PaymentAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"PaymentAccountRef"`
		PaymentAmount string `json:"PaymentAmount"`
		PaymentDate   string `json:"PaymentDate"`
		Sparse        string `json:"sparse"`
		Id            string `json:"Id"`
		MetaData      struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"TaxPayment"`
	Time time.Time `json:"time"`
}
