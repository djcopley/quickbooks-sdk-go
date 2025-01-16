package model

type QuickbooksEntityInfo struct {
	EntityName string
}

type QuickbooksEntity interface {
	GetEntityInfo() *QuickbooksEntityInfo
}
