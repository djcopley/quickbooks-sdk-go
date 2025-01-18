package batch

type Operation string

const (
	Create Operation = "create"
	Update Operation = "update"
	Delete Operation = "delete"
)

type Request struct {
	BatchItemRequest []ItemRequest `json:"BatchItemRequest,omitempty"`
}

type ItemRequest struct {
	BId         string    `json:"bId,omitempty"`
	OptionsData string    `json:"optionsData,omitempty"`
	Operation   Operation `json:"operation,omitempty"`
	Query       string    `json:"Query,omitempty"`
	Resource    string    `json:"Resource,omitempty"`
}

type Response struct {
	BatchItemResponse []ItemResponse `json:"BatchItemResponse,omitempty"`
}

type ItemResponse struct {
	BId           string `json:"bId,omitempty"`
	ResourceName  string `json:"resourceName,omitempty"`
	QueryResponse string `json:"QueryResponse,omitempty"`
}
