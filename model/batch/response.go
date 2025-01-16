package batch

type Response struct {
	BatchItemResponse []BatchItemResponse `json:"BatchItemResponse,omitempty"`
}

type BatchItemResponse struct {
	BId           string `json:"bId,omitempty"`
	ResourceName  string `json:"resourceName,omitempty"`
	QueryResponse string `json:"QueryResponse,omitempty"`
}
