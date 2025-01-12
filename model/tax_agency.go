package model

import "time"

type TaxAgency struct {
	Time      time.Time `json:"time"`
	TaxAgency struct {
		SyncToken             string `json:"SyncToken"`
		Domain                string `json:"domain"`
		DisplayName           string `json:"DisplayName"`
		TaxTrackedOnSales     bool   `json:"TaxTrackedOnSales"`
		TaxTrackedOnPurchases bool   `json:"TaxTrackedOnPurchases"`
		Sparse                bool   `json:"sparse"`
		Id                    string `json:"Id"`
		MetaData              struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"TaxAgency"`
}
