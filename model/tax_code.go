package model

import "time"

type TaxCode struct {
	TaxCode struct {
		SyncToken           string `json:"SyncToken"`
		Domain              string `json:"domain"`
		TaxGroup            bool   `json:"TaxGroup"`
		Name                string `json:"Name"`
		Taxable             bool   `json:"Taxable"`
		PurchaseTaxRateList struct {
			TaxRateDetail []interface{} `json:"TaxRateDetail"`
		} `json:"PurchaseTaxRateList"`
		Sparse      bool   `json:"sparse"`
		Active      bool   `json:"Active"`
		Description string `json:"Description"`
		MetaData    struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		Id               string `json:"Id"`
		SalesTaxRateList struct {
			TaxRateDetail []struct {
				TaxTypeApplicable string `json:"TaxTypeApplicable"`
				TaxRateRef        struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"TaxRateRef"`
				TaxOrder int `json:"TaxOrder"`
			} `json:"TaxRateDetail"`
		} `json:"SalesTaxRateList"`
	} `json:"TaxCode"`
	Time time.Time `json:"time"`
}
