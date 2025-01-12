package model

type TaxService struct {
	TaxRateDetails []struct {
		RateValue       string `json:"RateValue"`
		TaxRateId       string `json:"TaxRateId"`
		TaxApplicableOn string `json:"TaxApplicableOn"`
		TaxAgencyId     string `json:"TaxAgencyId"`
		TaxRateName     string `json:"TaxRateName"`
	} `json:"TaxRateDetails"`
	TaxCodeId string `json:"TaxCodeId"`
	TaxCode   string `json:"TaxCode"`
}
