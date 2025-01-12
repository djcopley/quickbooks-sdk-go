package model

import "time"

type InventoryAdjustment struct {
	InventoryAdjustment struct {
		DocNumber        string `json:"DocNumber"`
		TxnDate          string `json:"TxnDate"`
		PrivateNote      string `json:"PrivateNote"`
		SyncToken        string `json:"SyncToken"`
		Domain           string `json:"domain"`
		AdjustAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"AdjustAccountRef"`
		Sparse bool `json:"sparse"`
		Line   []struct {
			ItemAdjustmentLineDetail struct {
				QtyDiff int `json:"QtyDiff"`
				ItemRef struct {
					Value string `json:"value"`
				} `json:"ItemRef"`
			} `json:"ItemAdjustmentLineDetail"`
			Id string `json:"Id"`
		} `json:"Line"`
		Id       string `json:"Id"`
		MetaData struct {
			LastUpdatedTime string `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"InventoryAdjustment"`
	Time time.Time `json:"time"`
}
