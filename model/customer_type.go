package model

import "time"

type CustomerType struct {
	CustomerType struct {
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		Name      string `json:"Name"`
		Sparse    bool   `json:"sparse"`
		Active    bool   `json:"Active"`
		Id        string `json:"Id"`
		MetaData  struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"CustomerType"`
	Time time.Time `json:"time"`
}
