package model

import "time"

type FECReport struct {
	Header struct {
		ReportName string `json:"ReportName"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		ReportBasis string    `json:"ReportBasis"`
		StartPeriod string    `json:"StartPeriod"`
		Currency    string    `json:"Currency"`
		EndPeriod   string    `json:"EndPeriod"`
		Time        time.Time `json:"Time"`
	} `json:"Header"`
	Rows struct {
		Row []struct {
			ColData []struct {
				Value string `json:"value"`
				Id    string `json:"id,omitempty"`
			} `json:"ColData"`
			Type string `json:"type"`
		} `json:"Row"`
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColType  string `json:"ColType"`
			ColTitle string `json:"ColTitle"`
			MetaData []struct {
				Name  string `json:"Name"`
				Value string `json:"Value"`
			} `json:"MetaData"`
		} `json:"Column"`
	} `json:"Columns"`
}
