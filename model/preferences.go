package model

import "time"

type Preferences struct {
	Preferences struct {
		EmailMessagesPrefs struct {
			InvoiceMessage struct {
				Message string `json:"Message"`
				Subject string `json:"Subject"`
			} `json:"InvoiceMessage"`
			EstimateMessage struct {
				Message string `json:"Message"`
				Subject string `json:"Subject"`
			} `json:"EstimateMessage"`
			SalesReceiptMessage struct {
				Message string `json:"Message"`
				Subject string `json:"Subject"`
			} `json:"SalesReceiptMessage"`
			StatementMessage struct {
				Message string `json:"Message"`
				Subject string `json:"Subject"`
			} `json:"StatementMessage"`
		} `json:"EmailMessagesPrefs"`
		ProductAndServicesPrefs struct {
			QuantityWithPriceAndRate bool `json:"QuantityWithPriceAndRate"`
			ForPurchase              bool `json:"ForPurchase"`
			QuantityOnHand           bool `json:"QuantityOnHand"`
			ForSales                 bool `json:"ForSales"`
		} `json:"ProductAndServicesPrefs"`
		Domain      string `json:"domain"`
		SyncToken   string `json:"SyncToken"`
		ReportPrefs struct {
			ReportBasis                string `json:"ReportBasis"`
			CalcAgingReportFromTxnDate bool   `json:"CalcAgingReportFromTxnDate"`
		} `json:"ReportPrefs"`
		AccountingInfoPrefs struct {
			FirstMonthOfFiscalYear  string `json:"FirstMonthOfFiscalYear"`
			UseAccountNumbers       bool   `json:"UseAccountNumbers"`
			TaxYearMonth            string `json:"TaxYearMonth"`
			ClassTrackingPerTxn     bool   `json:"ClassTrackingPerTxn"`
			TrackDepartments        bool   `json:"TrackDepartments"`
			TaxForm                 string `json:"TaxForm"`
			CustomerTerminology     string `json:"CustomerTerminology"`
			BookCloseDate           string `json:"BookCloseDate"`
			DepartmentTerminology   string `json:"DepartmentTerminology"`
			ClassTrackingPerTxnLine bool   `json:"ClassTrackingPerTxnLine"`
		} `json:"AccountingInfoPrefs"`
		SalesFormsPrefs struct {
			ETransactionPaymentEnabled bool   `json:"ETransactionPaymentEnabled"`
			CustomTxnNumbers           bool   `json:"CustomTxnNumbers"`
			AllowShipping              bool   `json:"AllowShipping"`
			AllowServiceDate           bool   `json:"AllowServiceDate"`
			ETransactionEnabledStatus  string `json:"ETransactionEnabledStatus"`
			DefaultCustomerMessage     string `json:"DefaultCustomerMessage"`
			EmailCopyToCompany         bool   `json:"EmailCopyToCompany"`
			AllowEstimates             bool   `json:"AllowEstimates"`
			DefaultTerms               struct {
				Value string `json:"value"`
			} `json:"DefaultTerms"`
			AllowDiscount          bool   `json:"AllowDiscount"`
			DefaultDiscountAccount string `json:"DefaultDiscountAccount"`
			AllowDeposit           bool   `json:"AllowDeposit"`
			AutoApplyPayments      bool   `json:"AutoApplyPayments"`
			IPNSupportEnabled      bool   `json:"IPNSupportEnabled"`
			AutoApplyCredit        bool   `json:"AutoApplyCredit"`
			CustomField            []struct {
				CustomField []struct {
					BooleanValue bool   `json:"BooleanValue,omitempty"`
					Type         string `json:"Type"`
					Name         string `json:"Name"`
					StringValue  string `json:"StringValue,omitempty"`
				} `json:"CustomField"`
			} `json:"CustomField"`
			UsingPriceLevels      bool `json:"UsingPriceLevels"`
			ETransactionAttachPDF bool `json:"ETransactionAttachPDF"`
		} `json:"SalesFormsPrefs"`
		VendorAndPurchasesPrefs struct {
			BillableExpenseTracking bool `json:"BillableExpenseTracking"`
			TrackingByCustomer      bool `json:"TrackingByCustomer"`
			POCustomField           []struct {
				CustomField []struct {
					BooleanValue bool   `json:"BooleanValue,omitempty"`
					Type         string `json:"Type"`
					Name         string `json:"Name"`
					StringValue  string `json:"StringValue,omitempty"`
				} `json:"CustomField"`
			} `json:"POCustomField"`
		} `json:"VendorAndPurchasesPrefs"`
		TaxPrefs struct {
			TaxGroupCodeRef struct {
				Value string `json:"value"`
			} `json:"TaxGroupCodeRef"`
			UsingSalesTax bool `json:"UsingSalesTax"`
		} `json:"TaxPrefs"`
		OtherPrefs struct {
			NameValue []struct {
				Name  string `json:"Name"`
				Value string `json:"Value"`
			} `json:"NameValue"`
		} `json:"OtherPrefs"`
		Sparse            bool `json:"sparse"`
		TimeTrackingPrefs struct {
			WorkWeekStartDate       string `json:"WorkWeekStartDate"`
			MarkTimeEntriesBillable bool   `json:"MarkTimeEntriesBillable"`
			ShowBillRateToAll       bool   `json:"ShowBillRateToAll"`
			UseServices             bool   `json:"UseServices"`
			BillCustomers           bool   `json:"BillCustomers"`
		} `json:"TimeTrackingPrefs"`
		CurrencyPrefs struct {
			HomeCurrency struct {
				Value string `json:"value"`
			} `json:"HomeCurrency"`
			MultiCurrencyEnabled bool `json:"MultiCurrencyEnabled"`
		} `json:"CurrencyPrefs"`
		Id       string `json:"Id"`
		MetaData struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"Preferences"`
	Time time.Time `json:"time"`
}
