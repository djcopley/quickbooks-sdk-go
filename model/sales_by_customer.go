package model

import "time"

type SalesByCustomer struct {
	Header struct {
		Customer   string `json:"Customer"`
		ReportName string `json:"ReportName"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		ReportBasis        string    `json:"ReportBasis"`
		StartPeriod        string    `json:"StartPeriod"`
		Currency           string    `json:"Currency"`
		EndPeriod          string    `json:"EndPeriod"`
		Time               time.Time `json:"Time"`
		SummarizeColumnsBy string    `json:"SummarizeColumnsBy"`
	} `json:"Header"`
	Rows struct {
		Row []struct {
			ColData []struct {
				Id    string `json:"id,omitempty"`
				Value string `json:"value"`
			} `json:"ColData,omitempty"`
			Group   string `json:"group,omitempty"`
			Type    string `json:"type,omitempty"`
			Summary struct {
				ColData []struct {
					Value string `json:"value"`
				} `json:"ColData"`
			} `json:"Summary,omitempty"`
		} `json:"Row"`
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColType  string `json:"ColType"`
			ColTitle string `json:"ColTitle"`
		} `json:"Column"`
	} `json:"Columns"`
}
