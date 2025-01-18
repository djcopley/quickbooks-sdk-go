package bill

import "github.com/djcopley/quickbooks-sdk-go"

type Response struct {
	Bill Bill   `json:"Bill,omitempty"`
	Time string `json:"time,omitempty"`
}

type APAccountRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type VendorRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type CurrencyRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type LinkedTxn struct {
	TxnID   string `json:"TxnId,omitempty"`
	TxnType string `json:"TxnType,omitempty"`
}

type SalesTermRef struct {
	Value string `json:"value,omitempty"`
}

type ProjectRef struct {
	Value string `json:"value,omitempty"`
}

type TaxCodeRef struct {
	Value string `json:"value,omitempty"`
}

type AccountRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type CustomerRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type AccountBasedExpenseLineDetail struct {
	TaxCodeRef     TaxCodeRef  `json:"TaxCodeRef,omitempty"`
	AccountRef     AccountRef  `json:"AccountRef,omitempty"`
	BillableStatus string      `json:"BillableStatus,omitempty"`
	CustomerRef    CustomerRef `json:"CustomerRef,omitempty"`
}

type Line struct {
	Description                   string                        `json:"Description,omitempty"`
	DetailType                    string                        `json:"DetailType,omitempty"`
	ProjectRef                    ProjectRef                    `json:"ProjectRef,omitempty"`
	Amount                        float64                       `json:"Amount,omitempty"`
	ID                            string                        `json:"Id,omitempty"`
	AccountBasedExpenseLineDetail AccountBasedExpenseLineDetail `json:"AccountBasedExpenseLineDetail,omitempty"`
}

type MetaData struct {
	CreateTime      string `json:"CreateTime,omitempty"`
	LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`
}

type Bill struct {
	SyncToken    string       `json:"SyncToken,omitempty"`
	Domain       string       `json:"domain,omitempty"`
	APAccountRef APAccountRef `json:"APAccountRef,omitempty"`
	VendorRef    VendorRef    `json:"VendorRef,omitempty"`
	TxnDate      string       `json:"TxnDate,omitempty"`
	TotalAmt     float64      `json:"TotalAmt,omitempty"`
	CurrencyRef  CurrencyRef  `json:"CurrencyRef,omitempty"`
	LinkedTxn    []LinkedTxn  `json:"LinkedTxn,omitempty"`
	SalesTermRef SalesTermRef `json:"SalesTermRef,omitempty"`
	DueDate      string       `json:"DueDate,omitempty"`
	Sparse       bool         `json:"sparse,omitempty"`
	Line         []Line       `json:"Line,omitempty"`
	Balance      int          `json:"Balance,omitempty"`
	ID           string       `json:"Id,omitempty"`
	MetaData     MetaData     `json:"MetaData,omitempty"`
}

func (b *Bill) GetEntityInfo() *quickbooks.EntityInfo {
	return &quickbooks.EntityInfo{EntityName: "bill"}
}
