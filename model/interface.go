package model

type QuickbooksObjectInfo struct {
	EntityName string
}

type QuickbooksEntity interface {
	GetObjectInfo() *QuickbooksObjectInfo
}
