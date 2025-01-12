package model

import "time"

type Bill struct {
	Bill struct {
		SyncToken    string `json:"SyncToken"`
		Domain       string `json:"domain"`
		APAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"APAccountRef"`
		VendorRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"VendorRef"`
		TxnDate     string  `json:"TxnDate"`
		TotalAmt    float64 `json:"TotalAmt"`
		CurrencyRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CurrencyRef"`
		LinkedTxn []struct {
			TxnId   string `json:"TxnId"`
			TxnType string `json:"TxnType"`
		} `json:"LinkedTxn"`
		SalesTermRef struct {
			Value string `json:"value"`
		} `json:"SalesTermRef"`
		DueDate string `json:"DueDate"`
		Sparse  bool   `json:"sparse"`
		Line    []struct {
			Description string `json:"Description"`
			DetailType  string `json:"DetailType"`
			ProjectRef  struct {
				Value string `json:"value"`
			} `json:"ProjectRef"`
			Amount                        float64 `json:"Amount"`
			Id                            string  `json:"Id"`
			AccountBasedExpenseLineDetail struct {
				TaxCodeRef struct {
					Value string `json:"value"`
				} `json:"TaxCodeRef"`
				AccountRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"AccountRef"`
				BillableStatus string `json:"BillableStatus"`
				CustomerRef    struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"CustomerRef"`
			} `json:"AccountBasedExpenseLineDetail"`
		} `json:"Line"`
		Balance  int    `json:"Balance"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Bill"`
	Time time.Time `json:"time"`
}
