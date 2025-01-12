package model

import "time"

type BillPayment struct {
	BillPayment struct {
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		VendorRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"VendorRef"`
		TxnDate     string  `json:"TxnDate"`
		TotalAmt    float64 `json:"TotalAmt"`
		PayType     string  `json:"PayType"`
		PrivateNote string  `json:"PrivateNote"`
		Sparse      bool    `json:"sparse"`
		Line        []struct {
			Amount    float64 `json:"Amount"`
			LinkedTxn []struct {
				TxnId   string `json:"TxnId"`
				TxnType string `json:"TxnType"`
			} `json:"LinkedTxn"`
		} `json:"Line"`
		Id           string `json:"Id"`
		CheckPayment struct {
			PrintStatus    string `json:"PrintStatus"`
			BankAccountRef struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"BankAccountRef"`
		} `json:"CheckPayment"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"BillPayment"`
	Time time.Time `json:"time"`
}
