package batch

type Operation string

const (
	Create Operation = "create"
	Update Operation = "update"
	Delete Operation = "delete"
)

type Request struct {
	BatchItemRequest []BatchItemRequest `json:"BatchItemRequest,omitempty"`
}

type BatchItemRequest struct {
	BId          string    `json:"bId,omitempty"`
	OptionsData  string    `json:"optionsData,omitempty"`
	Operation    Operation `json:"operation,omitempty"`
	Query        string    `json:"Query,omitempty"`
	ResourceName string    `json:"resourceName,omitempty"`
}
