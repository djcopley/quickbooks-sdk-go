package model

import "time"

type Item struct {
	Item struct {
		FullyQualifiedName string `json:"FullyQualifiedName"`
		Domain             string `json:"domain"`
		Name               string `json:"Name"`
		SyncToken          string `json:"SyncToken"`
		Sparse             bool   `json:"sparse"`
		Active             bool   `json:"Active"`
		Type               string `json:"Type"`
		Id                 string `json:"Id"`
		MetaData           struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Item"`
	Time time.Time `json:"time"`
}
