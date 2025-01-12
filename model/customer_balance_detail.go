package model

import "time"

type CustomerBalanceDetail struct {
	Header struct {
		Customer   string `json:"Customer"`
		ReportName string `json:"ReportName"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		DateMacro string    `json:"DateMacro"`
		Currency  string    `json:"Currency"`
		Time      time.Time `json:"Time"`
	} `json:"Header"`
	Rows struct {
		Row []struct {
			Header struct {
				ColData []struct {
					Id    string `json:"id,omitempty"`
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
			Type    string `json:"type"`
			Summary struct {
				ColData []struct {
					Value string `json:"value"`
				} `json:"ColData"`
			} `json:"Summary"`
		} `json:"Row"`
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColType  string `json:"ColType"`
			ColTitle string `json:"ColTitle"`
		} `json:"Column"`
	} `json:"Columns"`
}
