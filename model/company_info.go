package model

import "time"

type CompanyInfo struct {
	CompanyInfo struct {
		SyncToken string `json:"SyncToken"`
		Domain    string `json:"domain"`
		LegalAddr struct {
			City                   string `json:"City"`
			Country                string `json:"Country"`
			Line1                  string `json:"Line1"`
			PostalCode             string `json:"PostalCode"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			Id                     string `json:"Id"`
		} `json:"LegalAddr"`
		SupportedLanguages string `json:"SupportedLanguages"`
		CompanyName        string `json:"CompanyName"`
		Country            string `json:"Country"`
		CompanyAddr        struct {
			City                   string `json:"City"`
			Country                string `json:"Country"`
			Line1                  string `json:"Line1"`
			PostalCode             string `json:"PostalCode"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			Id                     string `json:"Id"`
		} `json:"CompanyAddr"`
		Sparse  bool   `json:"sparse"`
		Id      string `json:"Id"`
		WebAddr struct {
		} `json:"WebAddr"`
		FiscalYearStartMonth      string `json:"FiscalYearStartMonth"`
		CustomerCommunicationAddr struct {
			City                   string `json:"City"`
			Country                string `json:"Country"`
			Line1                  string `json:"Line1"`
			PostalCode             string `json:"PostalCode"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			Id                     string `json:"Id"`
		} `json:"CustomerCommunicationAddr"`
		PrimaryPhone struct {
			FreeFormNumber string `json:"FreeFormNumber"`
		} `json:"PrimaryPhone"`
		LegalName        string `json:"LegalName"`
		CompanyStartDate string `json:"CompanyStartDate"`
		Email            struct {
			Address string `json:"Address"`
		} `json:"Email"`
		NameValue []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"NameValue"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"CompanyInfo"`
	Time time.Time `json:"time"`
}
