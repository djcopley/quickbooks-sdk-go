package model

import "time"

type Invoice struct {
	Invoice struct {
		TxnDate      string `json:"TxnDate"`
		Domain       string `json:"domain"`
		PrintStatus  string `json:"PrintStatus"`
		SalesTermRef struct {
			Value string `json:"value"`
		} `json:"SalesTermRef"`
		TotalAmt float64 `json:"TotalAmt"`
		Line     []struct {
			Description         string `json:"Description,omitempty"`
			DetailType          string `json:"DetailType"`
			SalesItemLineDetail struct {
				TaxCodeRef struct {
					Value string `json:"value"`
				} `json:"TaxCodeRef"`
				Qty       int     `json:"Qty"`
				UnitPrice float64 `json:"UnitPrice"`
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
		DueDate               string `json:"DueDate"`
		ApplyTaxAfterDiscount bool   `json:"ApplyTaxAfterDiscount"`
		DocNumber             string `json:"DocNumber"`
		Sparse                bool   `json:"sparse"`
		CustomerMemo          struct {
			Value string `json:"value"`
		} `json:"CustomerMemo"`
		ProjectRef struct {
			Value string `json:"value"`
		} `json:"ProjectRef"`
		Deposit     int     `json:"Deposit"`
		Balance     float64 `json:"Balance"`
		CustomerRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CustomerRef"`
		TxnTaxDetail struct {
			TxnTaxCodeRef struct {
				Value string `json:"value"`
			} `json:"TxnTaxCodeRef"`
			TotalTax float64 `json:"TotalTax"`
			TaxLine  []struct {
				DetailType    string  `json:"DetailType"`
				Amount        float64 `json:"Amount"`
				TaxLineDetail struct {
					NetAmountTaxable float64 `json:"NetAmountTaxable"`
					TaxPercent       int     `json:"TaxPercent"`
					TaxRateRef       struct {
						Value string `json:"value"`
					} `json:"TaxRateRef"`
					PercentBased bool `json:"PercentBased"`
				} `json:"TaxLineDetail"`
			} `json:"TaxLine"`
		} `json:"TxnTaxDetail"`
		SyncToken string `json:"SyncToken"`
		LinkedTxn []struct {
			TxnId   string `json:"TxnId"`
			TxnType string `json:"TxnType"`
		} `json:"LinkedTxn"`
		BillEmail struct {
			Address string `json:"Address"`
		} `json:"BillEmail"`
		ShipAddr struct {
			City                   string `json:"City"`
			Line1                  string `json:"Line1"`
			PostalCode             string `json:"PostalCode"`
			Lat                    string `json:"Lat"`
			Long                   string `json:"Long"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			Id                     string `json:"Id"`
		} `json:"ShipAddr"`
		EmailStatus string `json:"EmailStatus"`
		BillAddr    struct {
			Line4 string `json:"Line4"`
			Line3 string `json:"Line3"`
			Line2 string `json:"Line2"`
			Line1 string `json:"Line1"`
			Long  string `json:"Long"`
			Lat   string `json:"Lat"`
			Id    string `json:"Id"`
		} `json:"BillAddr"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		CustomField []struct {
			DefinitionId string `json:"DefinitionId"`
			StringValue  string `json:"StringValue"`
			Type         string `json:"Type"`
			Name         string `json:"Name"`
		} `json:"CustomField"`
		Id string `json:"Id"`
	} `json:"Invoice"`
	Time time.Time `json:"time"`
}
