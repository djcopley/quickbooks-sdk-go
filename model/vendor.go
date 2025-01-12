package model

import "time"

type Vendor struct {
	Vendor struct {
		PrimaryEmailAddr struct {
			Address string `json:"Address"`
		} `json:"PrimaryEmailAddr"`
		Vendor1099  bool   `json:"Vendor1099"`
		Domain      string `json:"domain"`
		GivenName   string `json:"GivenName"`
		DisplayName string `json:"DisplayName"`
		BillAddr    struct {
			City                   string `json:"City"`
			Line1                  string `json:"Line1"`
			PostalCode             string `json:"PostalCode"`
			Lat                    string `json:"Lat"`
			Long                   string `json:"Long"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			Id                     string `json:"Id"`
		} `json:"BillAddr"`
		SyncToken        string `json:"SyncToken"`
		PrintOnCheckName string `json:"PrintOnCheckName"`
		FamilyName       string `json:"FamilyName"`
		PrimaryPhone     struct {
			FreeFormNumber string `json:"FreeFormNumber"`
		} `json:"PrimaryPhone"`
		AcctNum     string `json:"AcctNum"`
		CompanyName string `json:"CompanyName"`
		WebAddr     struct {
			URI string `json:"URI"`
		} `json:"WebAddr"`
		Sparse   bool   `json:"sparse"`
		Active   bool   `json:"Active"`
		Balance  int    `json:"Balance"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Vendor"`
	Time time.Time `json:"time"`
}
