package model

import "time"

type Payment struct {
	Payment struct {
		SyncToken           string `json:"SyncToken"`
		Domain              string `json:"domain"`
		DepositToAccountRef struct {
			Value string `json:"value"`
		} `json:"DepositToAccountRef"`
		UnappliedAmt float64 `json:"UnappliedAmt"`
		TxnDate      string  `json:"TxnDate"`
		TotalAmt     float64 `json:"TotalAmt"`
		ProjectRef   struct {
			Value string `json:"value"`
		} `json:"ProjectRef"`
		ProcessPayment bool `json:"ProcessPayment"`
		Sparse         bool `json:"sparse"`
		Line           []struct {
			Amount float64 `json:"Amount"`
			LineEx struct {
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
			} `json:"LineEx"`
			LinkedTxn []struct {
				TxnId   string `json:"TxnId"`
				TxnType string `json:"TxnType"`
			} `json:"LinkedTxn"`
		} `json:"Line"`
		CustomerRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CustomerRef"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Payment"`
	Time time.Time `json:"time"`
}
