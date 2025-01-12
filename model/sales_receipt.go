package model

import "time"

type SalesReceipt struct {
	SalesReceipt struct {
		TxnDate       string  `json:"TxnDate"`
		Domain        string  `json:"domain"`
		PrintStatus   string  `json:"PrintStatus"`
		PaymentRefNum string  `json:"PaymentRefNum"`
		TotalAmt      float64 `json:"TotalAmt"`
		Line          []struct {
			Description         string `json:"Description,omitempty"`
			DetailType          string `json:"DetailType"`
			SalesItemLineDetail struct {
				TaxCodeRef struct {
					Value string `json:"value"`
				} `json:"TaxCodeRef"`
				Qty       float64 `json:"Qty"`
				UnitPrice int     `json:"UnitPrice"`
				ItemRef   struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"ItemRef"`
			} `json:"SalesItemLineDetail,omitempty"`
			LineNum            int     `json:"LineNum,omitempty"`
			Amount             float64 `json:"Amount"`
			Id                 string  `json:"Id,omitempty"`
			SubTotalLineDetail struct {
			} `json:"SubTotalLineDetail,omitempty"`
		} `json:"Line"`
		ApplyTaxAfterDiscount bool   `json:"ApplyTaxAfterDiscount"`
		DocNumber             string `json:"DocNumber"`
		Sparse                bool   `json:"sparse"`
		DepositToAccountRef   struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"DepositToAccountRef"`
		CustomerMemo struct {
			Value string `json:"value"`
		} `json:"CustomerMemo"`
		ProjectRef struct {
			Value string `json:"value"`
		} `json:"ProjectRef"`
		Balance     int `json:"Balance"`
		CustomerRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CustomerRef"`
		TxnTaxDetail struct {
			TotalTax int `json:"TotalTax"`
		} `json:"TxnTaxDetail"`
		SyncToken        string `json:"SyncToken"`
		PaymentMethodRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"PaymentMethodRef"`
		EmailStatus string `json:"EmailStatus"`
		BillAddr    struct {
			Lat   string `json:"Lat"`
			Long  string `json:"Long"`
			Id    string `json:"Id"`
			Line1 string `json:"Line1"`
		} `json:"BillAddr"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		CustomField []struct {
			DefinitionId string `json:"DefinitionId"`
			Type         string `json:"Type"`
			Name         string `json:"Name"`
		} `json:"CustomField"`
		Id string `json:"Id"`
	} `json:"SalesReceipt"`
	Time time.Time `json:"time"`
}
