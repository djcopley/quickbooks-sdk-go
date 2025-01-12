package model

import "time"

type Estimate struct {
	Estimate struct {
		DocNumber string `json:"DocNumber"`
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		TxnStatus string `json:"TxnStatus"`
		BillEmail struct {
			Address string `json:"Address"`
		} `json:"BillEmail"`
		TxnDate     string  `json:"TxnDate"`
		TotalAmt    float64 `json:"TotalAmt"`
		CustomerRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CustomerRef"`
		CustomerMemo struct {
			Value string `json:"value"`
		} `json:"CustomerMemo"`
		ShipAddr struct {
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			City                   string `json:"City"`
			PostalCode             string `json:"PostalCode"`
			Id                     string `json:"Id"`
			Line1                  string `json:"Line1"`
		} `json:"ShipAddr"`
		ProjectRef struct {
			Value string `json:"value"`
		} `json:"ProjectRef"`
		PrintStatus string `json:"PrintStatus"`
		BillAddr    struct {
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			City                   string `json:"City"`
			PostalCode             string `json:"PostalCode"`
			Id                     string `json:"Id"`
			Line1                  string `json:"Line1"`
		} `json:"BillAddr"`
		Sparse      bool   `json:"sparse"`
		EmailStatus string `json:"EmailStatus"`
		Line        []struct {
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
			DiscountLineDetail struct {
				DiscountAccountRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"DiscountAccountRef"`
				PercentBased    bool `json:"PercentBased"`
				DiscountPercent int  `json:"DiscountPercent"`
			} `json:"DiscountLineDetail,omitempty"`
		} `json:"Line"`
		ApplyTaxAfterDiscount bool `json:"ApplyTaxAfterDiscount"`
		CustomField           []struct {
			DefinitionId string `json:"DefinitionId"`
			Type         string `json:"Type"`
			Name         string `json:"Name"`
		} `json:"CustomField"`
		Id           string `json:"Id"`
		TxnTaxDetail struct {
			TotalTax int `json:"TotalTax"`
		} `json:"TxnTaxDetail"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Estimate"`
	Time time.Time `json:"time"`
}
