package quickbooks

type EntityInfo struct {
	EntityName string
}

type Entity interface {
	GetEntityInfo() *EntityInfo
}
