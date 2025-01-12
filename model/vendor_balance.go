package model

import "time"

type VendorBalance struct {
	Header struct {
		ReportName string `json:"ReportName"`
		DateMacro  string `json:"DateMacro"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		Currency string    `json:"Currency"`
		Time     time.Time `json:"Time"`
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