package model

import "time"

type JournalCode struct {
	JournalCode struct {
		SyncToken string    `json:"SyncToken"`
		Domain    string    `json:"domain"`
		Name      string    `json:"Name"`
		Sparse    bool      `json:"sparse"`
		Time      time.Time `json:"time"`
		Active    bool      `json:"Active"`
		MetaData  struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		Type        string `json:"Type"`
		Id          string `json:"Id"`
		Description string `json:"Description"`
	} `json:"JournalCode"`
}
