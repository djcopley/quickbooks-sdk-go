package model

import "time"

type VendorCredit struct {
	VendorCredit struct {
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		VendorRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"VendorRef"`
		TxnDate      string  `json:"TxnDate"`
		TotalAmt     float64 `json:"TotalAmt"`
		APAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"APAccountRef"`
		Sparse bool `json:"sparse"`
		Line   []struct {
			DetailType string  `json:"DetailType"`
			Amount     float64 `json:"Amount"`
			ProjectRef struct {
				Value string `json:"value"`
			} `json:"ProjectRef"`
			Id                            string `json:"Id"`
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
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"VendorCredit"`
	Time time.Time `json:"time"`
}
