package model

import "time"

type ARAgingDetail struct {
	Header struct {
		ReportName string `json:"ReportName"`
		Currency   string `json:"Currency"`
		EndPeriod  string `json:"EndPeriod"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		Time time.Time `json:"Time"`
	} `json:"Header"`
	Rows struct {
		Row []struct {
			Header struct {
				ColData []struct {
					Value string `json:"value"`
				} `json:"ColData"`
			} `json:"Header"`
			Rows struct {
				Row []struct {
					ColData []struct {
						Id    string `json:"id,omitempty"`
						Value string `json:"value"`
					} `json:"ColData"`
					Type string `json:"type"`
				} `json:"Row,omitempty"`
			} `json:"Rows,omitempty"`
			Type string `json:"type"`
		} `json:"Row"`
	} `json:"Rows"`
	Columns struct {
		Column []struct {
			ColType  string `json:"ColType"`
			ColTitle string `json:"ColTitle"`
		} `json:"Column"`
	} `json:"Columns"`
}
