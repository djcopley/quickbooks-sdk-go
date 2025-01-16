package model

type QuickbooksEntityInfo struct {
	EntityName string
	EntityType QuickbooksEntity
}

type QuickbooksEntity interface {
	GetEntityInfo() *QuickbooksEntityInfo
}
