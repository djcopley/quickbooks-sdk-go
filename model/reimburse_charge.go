package model

import "time"

type ReimburseCharge struct {
	ReimburseCharge struct {
		SyncToken       string `json:"SyncToken"`
		Domain          string `json:"domain"`
		HasBeenInvoiced bool   `json:"HasBeenInvoiced"`
		TxnDate         string `json:"TxnDate"`
		CurrencyRef     struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CurrencyRef"`
		LinkedTxn []struct {
			TxnId   string `json:"TxnId"`
			TxnType string `json:"TxnType"`
		} `json:"LinkedTxn"`
		Amount float64 `json:"Amount"`
		Sparse bool    `json:"sparse"`
		Line   []struct {
			LinkedTxn []struct {
				TxnId   string `json:"TxnId"`
				TxnType string `json:"TxnType"`
			} `json:"LinkedTxn"`
			DetailType          string `json:"DetailType"`
			ReimburseLineDetail struct {
				ItemRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"ItemRef,omitempty"`
				Qty        int `json:"Qty,omitempty"`
				TaxCodeRef struct {
					Value string `json:"value"`
				} `json:"TaxCodeRef"`
				MarkupInfo struct {
					Percent int `json:"Percent"`
				} `json:"MarkupInfo"`
				ItemAccountRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"ItemAccountRef"`
				UnitPrice int `json:"UnitPrice,omitempty"`
			} `json:"ReimburseLineDetail"`
			LineNum     int     `json:"LineNum"`
			Amount      float64 `json:"Amount"`
			Id          string  `json:"Id"`
			Description string  `json:"Description,omitempty"`
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
	} `json:"ReimburseCharge"`
	Time time.Time `json:"time"`
}
