package model

import "time"

type Class struct {
	Class struct {
		FullyQualifiedName string `json:"FullyQualifiedName"`
		Domain             string `json:"domain"`
		Name               string `json:"Name"`
		SyncToken          string `json:"SyncToken"`
		SubClass           bool   `json:"SubClass"`
		Sparse             bool   `json:"sparse"`
		Active             bool   `json:"Active"`
		Id                 string `json:"Id"`
		MetaData           struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Class"`
	Time time.Time `json:"time"`
}
