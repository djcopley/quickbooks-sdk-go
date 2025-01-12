package model

import "time"

type CompanyCurrency struct {
	CompanyCurrency struct {
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		Code      string `json:"Code"`
		Name      string `json:"Name"`
		Sparse    bool   `json:"sparse"`
		Active    bool   `json:"Active"`
		Id        string `json:"Id"`
		MetaData  struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"CompanyCurrency"`
	Time time.Time `json:"time"`
}
