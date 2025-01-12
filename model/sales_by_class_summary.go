package model

import "time"

type SalesByClassSummary struct {
	Header struct {
		ReportName string `json:"ReportName"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		DateMacro          string    `json:"DateMacro"`
		ReportBasis        string    `json:"ReportBasis"`
		StartPeriod        string    `json:"StartPeriod"`
		Currency           string    `json:"Currency"`
		EndPeriod          string    `json:"EndPeriod"`
		Time               time.Time `json:"Time"`
		SummarizeColumnsBy string    `json:"SummarizeColumnsBy"`
	} `json:"Header"`
	Rows struct {
		Row []struct {
			Header struct {
				ColData []struct {
					Value string `json:"value"`
				} `json:"ColData"`
			} `json:"Header,omitempty"`
			Rows struct {
				Row []struct {
					ColData []struct {
						Value string `json:"value"`
					} `json:"ColData"`
					Type string `json:"type"`
				} `json:"Row"`
			} `json:"Rows,omitempty"`
			Type    string `json:"type,omitempty"`
			Group   string `json:"group"`
			Summary struct {
				ColData []struct {
					Value string `json:"value"`
				} `json:"ColData"`
			} `json:"Summary,omitempty"`
			ColData []struct {
				Value string `json:"value"`
			} `json:"ColData,omitempty"`
		} `json:"Row"`
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColType  string `json:"ColType"`
			ColTitle string `json:"ColTitle"`
		} `json:"Column"`
	} `json:"Columns"`
}
