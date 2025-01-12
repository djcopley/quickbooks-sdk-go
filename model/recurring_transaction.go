package model

import "time"

type RecurringTransaction struct {
	QueryResponse struct {
		StartPosition        int `json:"startPosition"`
		MaxResults           int `json:"maxResults"`
		RecurringTransaction []struct {
			Bill struct {
				SyncToken     string `json:"SyncToken"`
				Domain        string `json:"domain"`
				RecurringInfo struct {
					Active       bool   `json:"Active"`
					RecurType    string `json:"RecurType"`
					ScheduleInfo struct {
						NumInterval  int    `json:"NumInterval"`
						NextDate     string `json:"NextDate"`
						DayOfMonth   int    `json:"DayOfMonth"`
						PreviousDate string `json:"PreviousDate"`
						IntervalType string `json:"IntervalType"`
					} `json:"ScheduleInfo"`
					Name string `json:"Name"`
				} `json:"RecurringInfo"`
				RecurDataRef struct {
					Value string `json:"value"`
				} `json:"RecurDataRef"`
				CurrencyRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"CurrencyRef"`
				TotalAmt     float64 `json:"TotalAmt"`
				APAccountRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"APAccountRef"`
				Id        string `json:"Id"`
				Sparse    bool   `json:"sparse"`
				VendorRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"VendorRef"`
				Line []struct {
					Description                   string  `json:"Description"`
					DetailType                    string  `json:"DetailType"`
					LineNum                       int     `json:"LineNum"`
					Amount                        float64 `json:"Amount"`
					Id                            string  `json:"Id"`
					AccountBasedExpenseLineDetail struct {
						TaxCodeRef struct {
							Value string `json:"value"`
						} `json:"TaxCodeRef"`
						AccountRef struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"AccountRef"`
						BillableStatus string `json:"BillableStatus"`
					} `json:"AccountBasedExpenseLineDetail"`
				} `json:"Line"`
				Balance      float64 `json:"Balance"`
				SalesTermRef struct {
					Value string `json:"value"`
				} `json:"SalesTermRef"`
				MetaData struct {
					CreateTime      time.Time `json:"CreateTime"`
					LastUpdatedTime time.Time `json:"LastUpdatedTime"`
				} `json:"MetaData"`
			} `json:"Bill"`
		} `json:"RecurringTransaction"`
	} `json:"QueryResponse"`
	Time time.Time `json:"time"`
}
