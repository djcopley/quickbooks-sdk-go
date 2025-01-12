package model

import "time"

type Transfer struct {
	Transfer struct {
		SyncToken    string `json:"SyncToken"`
		Domain       string `json:"domain"`
		TxnDate      string `json:"TxnDate"`
		ToAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"ToAccountRef"`
		Amount         float64 `json:"Amount"`
		Sparse         bool    `json:"sparse"`
		Id             string  `json:"Id"`
		FromAccountRef struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"FromAccountRef"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Transfer"`
	Time time.Time `json:"time"`
}
