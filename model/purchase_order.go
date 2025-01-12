package model

import "time"

type PurchaseOrder struct {
	PurchaseOrder struct {
		DocNumber string `json:"DocNumber"`
		SyncToken string `json:"SyncToken"`
		POEmail   struct {
			Address string `json:"Address"`
		} `json:"POEmail"`
		APAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"APAccountRef"`
		CurrencyRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CurrencyRef"`
		TxnDate  string  `json:"TxnDate"`
		TotalAmt float64 `json:"TotalAmt"`
		ShipAddr struct {
			Line4 string `json:"Line4"`
			Line3 string `json:"Line3"`
			Id    string `json:"Id"`
			Line1 string `json:"Line1"`
			Line2 string `json:"Line2"`
		} `json:"ShipAddr"`
		Domain      string `json:"domain"`
		Id          string `json:"Id"`
		POStatus    string `json:"POStatus"`
		Sparse      bool   `json:"sparse"`
		EmailStatus string `json:"EmailStatus"`
		VendorRef   struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"VendorRef"`
		Line []struct {
			DetailType string  `json:"DetailType"`
			Amount     float64 `json:"Amount"`
			ProjectRef struct {
				Value string `json:"value"`
			} `json:"ProjectRef"`
			Id                         string `json:"Id"`
			ItemBasedExpenseLineDetail struct {
				ItemRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"ItemRef"`
				CustomerRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"CustomerRef"`
				Qty        int `json:"Qty"`
				TaxCodeRef struct {
					Value string `json:"value"`
				} `json:"TaxCodeRef"`
				BillableStatus string `json:"BillableStatus"`
				UnitPrice      int    `json:"UnitPrice"`
			} `json:"ItemBasedExpenseLineDetail"`
		} `json:"Line"`
		CustomField []struct {
			DefinitionId string `json:"DefinitionId"`
			Type         string `json:"Type"`
			Name         string `json:"Name"`
		} `json:"CustomField"`
		VendorAddr struct {
			Line4 string `json:"Line4"`
			Line3 string `json:"Line3"`
			Id    string `json:"Id"`
			Line1 string `json:"Line1"`
			Line2 string `json:"Line2"`
		} `json:"VendorAddr"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"PurchaseOrder"`
	Time time.Time `json:"time"`
}
