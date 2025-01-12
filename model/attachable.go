package model

import "time"

type Attachable struct {
	Attachable struct {
		SyncToken     string `json:"SyncToken"`
		Domain        string `json:"domain"`
		AttachableRef []struct {
			IncludeOnSend bool `json:"IncludeOnSend"`
			EntityRef     struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"EntityRef"`
		} `json:"AttachableRef"`
		Note     string `json:"Note"`
		Sparse   bool   `json:"sparse"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Attachable"`
	Time time.Time `json:"time"`
}
