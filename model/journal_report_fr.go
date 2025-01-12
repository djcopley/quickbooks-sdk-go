package model

import "time"

type JournalReportFR struct {
	Header struct {
		ReportName string `json:"ReportName"`
		Option     []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Option"`
		DateMacro   string    `json:"DateMacro"`
		StartPeriod string    `json:"StartPeriod"`
		Currency    string    `json:"Currency"`
		EndPeriod   string    `json:"EndPeriod"`
		Time        time.Time `json:"Time"`
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
						} `json:"Row"`
					} `json:"Rows"`
					Type    string `json:"type"`
					Summary struct {
						ColData []struct {
							Value string `json:"value"`
						} `json:"ColData"`
					} `json:"Summary"`
				} `json:"Row"`
			} `json:"Rows"`
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
