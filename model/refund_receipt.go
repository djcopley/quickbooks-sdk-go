package model

import "time"

type RefundReceipt struct {
	RefundReceipt struct {
		DocNumber        string `json:"DocNumber"`
		SyncToken        string `json:"SyncToken"`
		Domain           string `json:"domain"`
		Balance          int    `json:"Balance"`
		PaymentMethodRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"PaymentMethodRef"`
		BillAddr struct {
			Line4 string `json:"Line4"`
			Line3 string `json:"Line3"`
			Line2 string `json:"Line2"`
			Line1 string `json:"Line1"`
			Long  string `json:"Long"`
			Lat   string `json:"Lat"`
			Id    string `json:"Id"`
		} `json:"BillAddr"`
		DepositToAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"DepositToAccountRef"`
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
			City                   string `json:"City"`
			Line1                  string `json:"Line1"`
			PostalCode             string `json:"PostalCode"`
			Lat                    string `json:"Lat"`
			Long                   string `json:"Long"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			Id                     string `json:"Id"`
		} `json:"ShipAddr"`
		PrintStatus string `json:"PrintStatus"`
		BillEmail   struct {
			Address string `json:"Address"`
		} `json:"BillEmail"`
		Sparse bool `json:"sparse"`
		Line   []struct {
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
	} `json:"RefundReceipt"`
	Time time.Time `json:"time"`
}
