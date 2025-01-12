package model

import "time"

type Purchase struct {
	Purchase struct {
		SyncToken  string `json:"SyncToken"`
		Domain     string `json:"domain"`
		PurchaseEx struct {
			Any []struct {
				Name  string `json:"name"`
				Nil   bool   `json:"nil"`
				Value struct {
					Name  string `json:"Name"`
					Value string `json:"Value"`
				} `json:"value"`
				DeclaredType    string `json:"declaredType"`
				Scope           string `json:"scope"`
				GlobalScope     bool   `json:"globalScope"`
				TypeSubstituted bool   `json:"typeSubstituted"`
			} `json:"any"`
		} `json:"PurchaseEx"`
		TxnDate     string  `json:"TxnDate"`
		TotalAmt    float64 `json:"TotalAmt"`
		PaymentType string  `json:"PaymentType"`
		Sparse      bool    `json:"sparse"`
		Line        []struct {
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
			} `json:"AccountBasedExpenseLineDetail"`
		} `json:"Line"`
		AccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"AccountRef"`
		CustomField []interface{} `json:"CustomField"`
		Id          string        `json:"Id"`
		MetaData    struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Purchase"`
	Time time.Time `json:"time"`
}
