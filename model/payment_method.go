package model

import "time"

type PaymentMethod struct {
	PaymentMethod struct {
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		Name      string `json:"Name"`
		Sparse    bool   `json:"sparse"`
		Active    bool   `json:"Active"`
		Type      string `json:"Type"`
		Id        string `json:"Id"`
		MetaData  struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"PaymentMethod"`
	Time time.Time `json:"time"`
}
