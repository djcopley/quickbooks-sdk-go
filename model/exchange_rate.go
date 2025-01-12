package model

import "time"

type ExchangeRate struct {
	ExchangeRate struct {
		SyncToken          string  `json:"SyncToken"`
		Domain             string  `json:"domain"`
		AsOfDate           string  `json:"AsOfDate"`
		SourceCurrencyCode string  `json:"SourceCurrencyCode"`
		Rate               float64 `json:"Rate"`
		Sparse             bool    `json:"sparse"`
		TargetCurrencyCode string  `json:"TargetCurrencyCode"`
		MetaData           struct {
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"ExchangeRate"`
	Time time.Time `json:"time"`
}
