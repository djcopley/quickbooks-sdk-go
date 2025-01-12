package model

import "time"

type SalesByProduct struct {
	Header struct {
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
			} `json:"ColData"`
			Group string `json:"group,omitempty"`
		} `json:"Row"`
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColType  string `json:"ColType"`
			ColTitle string `json:"ColTitle"`
			Columns  struct {
				Column []struct {
					ColType  string `json:"ColType"`
					ColTitle string `json:"ColTitle"`
					MetaData []struct {
						Name  string `json:"Name"`
						Value string `json:"Value"`
					} `json:"MetaData,omitempty"`
				} `json:"Column"`
			} `json:"Columns,omitempty"`
		} `json:"Column"`
	} `json:"Columns"`
}
