package model

import "time"

type Budget struct {
	QueryResponse struct {
		StartPosition int `json:"startPosition"`
		TotalCount    int `json:"totalCount"`
		Budget        []struct {
			StartDate       string `json:"StartDate"`
			BudgetEntryType string `json:"BudgetEntryType"`
			EndDate         string `json:"EndDate"`
			Name            string `json:"Name"`
			SyncToken       string `json:"SyncToken"`
			BudgetType      string `json:"BudgetType"`
			Domain          string `json:"domain"`
			Sparse          bool   `json:"sparse"`
			Active          bool   `json:"Active"`
			BudgetDetail    []struct {
				Amount     float64 `json:"Amount"`
				AccountRef struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"AccountRef"`
				BudgetDate string `json:"BudgetDate"`
			} `json:"BudgetDetail"`
			Id       string `json:"Id"`
			MetaData struct {
				CreateTime      time.Time `json:"CreateTime"`
				LastUpdatedTime time.Time `json:"LastUpdatedTime"`
			} `json:"MetaData"`
		} `json:"Budget"`
		MaxResults int `json:"maxResults"`
	} `json:"QueryResponse"`
	Time time.Time `json:"time"`
}
