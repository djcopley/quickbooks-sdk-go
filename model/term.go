package model

import "time"

type Term struct {
	Term struct {
		SyncToken       string `json:"SyncToken"`
		Domain          string `json:"domain"`
		Name            string `json:"Name"`
		DiscountPercent int    `json:"DiscountPercent"`
		DiscountDays    int    `json:"DiscountDays"`
		Type            string `json:"Type"`
		Sparse          bool   `json:"sparse"`
		Active          bool   `json:"Active"`
		DueDays         int    `json:"DueDays"`
		Id              string `json:"Id"`
		MetaData        struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Term"`
	Time time.Time `json:"time"`
}
