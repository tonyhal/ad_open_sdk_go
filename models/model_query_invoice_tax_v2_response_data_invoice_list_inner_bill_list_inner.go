/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInner struct for QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInner
type QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInner struct {
	//
	BillDetailList []*QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInnerBillDetailListInner `json:"bill_detail_list,omitempty"`
	//
	InvoiceCode *string `json:"invoice_code,omitempty"`
	//
	InvoiceDate *string `json:"invoice_date,omitempty"`
	//
	InvoiceNo *string `json:"invoice_no,omitempty"`
	//
	TotalPriceTax *float64 `json:"total_price_tax,omitempty"`
}