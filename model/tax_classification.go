package model

type TaxClassification struct {
	TaxClassification struct {
		ApplicableTo []string `json:"applicableTo"`
		Code         string   `json:"code"`
		Description  string   `json:"description"`
		Level        string   `json:"level"`
		ParentRef    struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"ParentRef"`
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"TaxClassification"`
}
