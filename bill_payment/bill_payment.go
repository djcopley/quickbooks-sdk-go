package bill_payment

import (
	"github.com/djcopley/quickbooks-sdk-go"
	"time"
)

type Response struct {
	BillPayment BillPayment `json:"BillPayment,omitempty"`
	Time        time.Time   `json:"time,omitempty"`
}

type VendorRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type LinkedTxn struct {
	TxnID   string `json:"TxnId,omitempty"`
	TxnType string `json:"TxnType,omitempty"`
}

type Line struct {
	Amount    float64     `json:"Amount,omitempty"`
	LinkedTxn []LinkedTxn `json:"LinkedTxn,omitempty"`
}

type BankAccountRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type CheckPayment struct {
	PrintStatus    string         `json:"PrintStatus,omitempty"`
	BankAccountRef BankAccountRef `json:"BankAccountRef,omitempty"`
}

type MetaData struct {
	CreateTime      time.Time `json:"CreateTime,omitempty"`
	LastUpdatedTime time.Time `json:"LastUpdatedTime,omitempty"`
}

type BillPayment struct {
	SyncToken    string       `json:"SyncToken,omitempty"`
	Domain       string       `json:"domain,omitempty"`
	VendorRef    VendorRef    `json:"VendorRef,omitempty"`
	TxnDate      string       `json:"TxnDate,omitempty"`
	TotalAmt     float64      `json:"TotalAmt,omitempty"`
	PayType      string       `json:"PayType,omitempty"`
	PrivateNote  string       `json:"PrivateNote,omitempty"`
	Sparse       bool         `json:"sparse,omitempty"`
	Line         []Line       `json:"Line,omitempty"`
	ID           string       `json:"Id,omitempty"`
	CheckPayment CheckPayment `json:"CheckPayment,omitempty"`
	MetaData     MetaData     `json:"MetaData,omitempty"`
}

func (b *BillPayment) GetEntityInfo() *quickbooks.EntityInfo {
	return &quickbooks.EntityInfo{EntityName: "billpayment"}
}
