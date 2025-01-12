package model

import "time"

type Employee struct {
	Employee struct {
		SyncToken    string `json:"SyncToken"`
		Domain       string `json:"domain"`
		DisplayName  string `json:"DisplayName"`
		PrimaryPhone struct {
			FreeFormNumber string `json:"FreeFormNumber"`
		} `json:"PrimaryPhone"`
		PrintOnCheckName string `json:"PrintOnCheckName"`
		FamilyName       string `json:"FamilyName"`
		Active           bool   `json:"Active"`
		SSN              string `json:"SSN"`
		PrimaryAddr      struct {
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			City                   string `json:"City"`
			PostalCode             string `json:"PostalCode"`
			Id                     string `json:"Id"`
			Line1                  string `json:"Line1"`
		} `json:"PrimaryAddr"`
		Sparse       bool   `json:"sparse"`
		BillableTime bool   `json:"BillableTime"`
		GivenName    string `json:"GivenName"`
		Id           string `json:"Id"`
		MetaData     struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Employee"`
	Time time.Time `json:"time"`
}
