package model

import "time"

type TaxRate struct {
	TaxRate struct {
		RateValue int `json:"RateValue"`
		AgencyRef struct {
			Value string `json:"value"`
		} `json:"AgencyRef"`
		Domain         string `json:"domain"`
		Name           string `json:"Name"`
		SyncToken      string `json:"SyncToken"`
		SpecialTaxType string `json:"SpecialTaxType"`
		DisplayType    string `json:"DisplayType"`
		Sparse         bool   `json:"sparse"`
		Active         bool   `json:"Active"`
		MetaData       struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		Id          string `json:"Id"`
		Description string `json:"Description"`
	} `json:"TaxRate"`
	Time time.Time `json:"time"`
}
