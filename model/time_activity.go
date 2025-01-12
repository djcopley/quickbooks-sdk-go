package model

import "time"

type TimeActivity struct {
	TimeActivity struct {
		TxnDate     string `json:"TxnDate"`
		Domain      string `json:"domain"`
		NameOf      string `json:"NameOf"`
		Description string `json:"Description"`
		ItemRef     struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"ItemRef"`
		Minutes    int `json:"Minutes"`
		ProjectRef struct {
			Value string `json:"value"`
		} `json:"ProjectRef"`
		Hours          int    `json:"Hours"`
		BillableStatus string `json:"BillableStatus"`
		Sparse         bool   `json:"sparse"`
		HourlyRate     int    `json:"HourlyRate"`
		Taxable        bool   `json:"Taxable"`
		EmployeeRef    struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"EmployeeRef"`
		SyncToken   string `json:"SyncToken"`
		CustomerRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CustomerRef"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"TimeActivity"`
	Time time.Time `json:"time"`
}
