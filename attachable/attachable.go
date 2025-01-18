package attachable

import "github.com/djcopley/quickbooks-sdk-go"

type Response struct {
	Attachable Attachable `json:"Attachable,omitempty"`
	Time       string     `json:"time,omitempty"`
}

type EntityRef struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Ref struct {
	IncludeOnSend bool      `json:"IncludeOnSend,omitempty"`
	EntityRef     EntityRef `json:"EntityRef,omitempty"`
}

type MetaData struct {
	CreateTime      string `json:"CreateTime,omitempty"`
	LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`
}

type Attachable struct {
	ID            string   `json:"Id,omitempty"`
	SyncToken     string   `json:"SyncToken,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	AttachableRef []Ref    `json:"AttachableRef,omitempty"`
	Note          string   `json:"Note,omitempty"`
	Sparse        bool     `json:"sparse,omitempty"`
	MetaData      MetaData `json:"MetaData,omitempty"`
}

func (a *Attachable) GetEntityInfo() *quickbooks.EntityInfo {
	return &quickbooks.EntityInfo{EntityName: "attachable"}
}
