package model

import "time"

type CreditMemo struct {
	CreditMemo struct {
		TxnDate         string  `json:"TxnDate"`
		Domain          string  `json:"domain"`
		PrintStatus     string  `json:"PrintStatus"`
		TotalAmt        float64 `json:"TotalAmt"`
		RemainingCredit int     `json:"RemainingCredit"`
		Line            []struct {
			Description         string `json:"Description,omitempty"`
			DetailType          string `json:"DetailType"`
			SalesItemLineDetail struct {
				TaxCodeRef struct {
					Value string `json:"value"`
				} `json:"TaxCodeRef"`
				Qty       int `json:"Qty"`
				UnitPrice int `json:"UnitPrice"`
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
		CustomerMemo          struct {
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
		SyncToken   string `json:"SyncToken"`
		CustomField []struct {
			DefinitionId string `json:"DefinitionId"`
			Type         string `json:"Type"`
			Name         string `json:"Name"`
		} `json:"CustomField"`
		ShipAddr struct {
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			City                   string `json:"City"`
			PostalCode             string `json:"PostalCode"`
			Id                     string `json:"Id"`
			Line1                  string `json:"Line1"`
		} `json:"ShipAddr"`
		EmailStatus string `json:"EmailStatus"`
		BillAddr    struct {
			Line4 string `json:"Line4"`
			Line3 string `json:"Line3"`
			Id    string `json:"Id"`
			Line1 string `json:"Line1"`
			Line2 string `json:"Line2"`
		} `json:"BillAddr"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		BillEmail struct {
			Address string `json:"Address"`
		} `json:"BillEmail"`
		Id string `json:"Id"`
	} `json:"CreditMemo"`
	Time time.Time `json:"time"`
}
