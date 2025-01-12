package model

import "time"

type TransactionListByVendor struct {
	Header struct {
		ReportName string `json:"ReportName"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		StartPeriod string    `json:"StartPeriod"`
		Currency    string    `json:"Currency"`
		EndPeriod   string    `json:"EndPeriod"`
		Time        time.Time `json:"Time"`
	} `json:"Header"`
	Rows struct {
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColTitle string `json:"ColTitle"`
			MetaData []struct {
				Name  string `json:"Name"`
				Value string `json:"Value"`
			} `json:"MetaData"`
		} `json:"Column"`
	} `json:"Columns"`
}
