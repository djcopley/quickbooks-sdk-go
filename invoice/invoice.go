package invoice

import (
	"github.com/djcopley/quickbooks-sdk-go"
	"time"
)

type Response struct {
	Invoice Invoice   `json:"Invoice,omitempty"`
	Time    time.Time `json:"time,omitempty"`
}

type SalesTermRef struct {
	Value string `json:"value,omitempty"`
}

type TaxCodeRef struct {
	Value string `json:"value,omitempty"`
}

type ItemRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type SalesItemLineDetail struct {
	TaxCodeRef TaxCodeRef `json:"TaxCodeRef,omitempty"`
	Qty        float64    `json:"Qty,omitempty"`
	UnitPrice  float64    `json:"UnitPrice,omitempty"`
	ItemRef    ItemRef    `json:"ItemRef,omitempty"`
}

type SubTotalLineDetail struct{}

type Line struct {
	Description         string              `json:"Description,omitempty"`
	DetailType          string              `json:"DetailType,omitempty"`
	SalesItemLineDetail SalesItemLineDetail `json:"SalesItemLineDetail,omitempty"`
	LineNum             int                 `json:"LineNum,omitempty"`
	Amount              float64             `json:"Amount,omitempty"`
	ID                  string              `json:"Id,omitempty"`
	SubTotalLineDetail  SubTotalLineDetail  `json:"SubTotalLineDetail,omitempty"`
}

type CustomerMemo struct {
	Value string `json:"value,omitempty"`
}

type ProjectRef struct {
	Value string `json:"value,omitempty"`
}

type CustomerRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TxnTaxCodeRef struct {
	Value string `json:"value,omitempty"`
}

type TaxRateRef struct {
	Value string `json:"value,omitempty"`
}

type TaxLineDetail struct {
	NetAmountTaxable float64    `json:"NetAmountTaxable,omitempty"`
	TaxPercent       float64    `json:"TaxPercent,omitempty"`
	TaxRateRef       TaxRateRef `json:"TaxRateRef,omitempty"`
	PercentBased     bool       `json:"PercentBased,omitempty"`
}

type TaxLine struct {
	DetailType    string        `json:"DetailType,omitempty"`
	Amount        float64       `json:"Amount,omitempty"`
	TaxLineDetail TaxLineDetail `json:"TaxLineDetail,omitempty"`
}

type TxnTaxDetail struct {
	TxnTaxCodeRef TxnTaxCodeRef `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      float64       `json:"TotalTax,omitempty"`
	TaxLine       []TaxLine     `json:"TaxLine,omitempty"`
}

type LinkedTxn struct {
	TxnID   string `json:"TxnId,omitempty"`
	TxnType string `json:"TxnType,omitempty"`
}

type BillEmail struct {
	Address string `json:"Address,omitempty"`
}

type ShipAddr struct {
	City                   string `json:"City,omitempty"`
	Line1                  string `json:"Line1,omitempty"`
	PostalCode             string `json:"PostalCode,omitempty"`
	Lat                    string `json:"Lat,omitempty"`
	Long                   string `json:"Long,omitempty"`
	CountrySubDivisionCode string `json:"CountrySubDivisionCode,omitempty"`
	ID                     string `json:"Id,omitempty"`
}

type BillAddr struct {
	Line4 string `json:"Line4,omitempty"`
	Line3 string `json:"Line3,omitempty"`
	Line2 string `json:"Line2,omitempty"`
	Line1 string `json:"Line1,omitempty"`
	Long  string `json:"Long,omitempty"`
	Lat   string `json:"Lat,omitempty"`
	ID    string `json:"Id,omitempty"`
}

type MetaData struct {
	CreateTime      string `json:"CreateTime,omitempty"`
	LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`
}

type CustomField struct {
	DefinitionID string `json:"DefinitionId,omitempty"`
	StringValue  string `json:"StringValue,omitempty"`
	Type         string `json:"Type,omitempty"`
	Name         string `json:"Name,omitempty"`
}

type Invoice struct {
	TxnDate               string        `json:"TxnDate,omitempty"`
	Domain                string        `json:"domain,omitempty"`
	PrintStatus           string        `json:"PrintStatus,omitempty"`
	SalesTermRef          SalesTermRef  `json:"SalesTermRef,omitempty"`
	TotalAmt              float64       `json:"TotalAmt,omitempty"`
	Line                  []Line        `json:"Line,omitempty"`
	DueDate               string        `json:"DueDate,omitempty"`
	ApplyTaxAfterDiscount bool          `json:"ApplyTaxAfterDiscount,omitempty"`
	DocNumber             string        `json:"DocNumber,omitempty"`
	Sparse                bool          `json:"sparse,omitempty"`
	CustomerMemo          CustomerMemo  `json:"CustomerMemo,omitempty"`
	ProjectRef            ProjectRef    `json:"ProjectRef,omitempty"`
	Deposit               int           `json:"Deposit,omitempty"`
	Balance               float64       `json:"Balance,omitempty"`
	CustomerRef           CustomerRef   `json:"CustomerRef,omitempty"`
	TxnTaxDetail          TxnTaxDetail  `json:"TxnTaxDetail,omitempty"`
	SyncToken             string        `json:"SyncToken,omitempty"`
	LinkedTxn             []LinkedTxn   `json:"LinkedTxn,omitempty"`
	BillEmail             BillEmail     `json:"BillEmail,omitempty"`
	ShipAddr              ShipAddr      `json:"ShipAddr,omitempty"`
	EmailStatus           string        `json:"EmailStatus,omitempty"`
	BillAddr              BillAddr      `json:"BillAddr,omitempty"`
	MetaData              MetaData      `json:"MetaData,omitempty"`
	CustomField           []CustomField `json:"CustomField,omitempty"`
	ID                    string        `json:"Id,omitempty"`
}

func (i *Invoice) GetEntityInfo() *quickbooks.EntityInfo {
	return &quickbooks.EntityInfo{EntityName: "invoice"}
}
