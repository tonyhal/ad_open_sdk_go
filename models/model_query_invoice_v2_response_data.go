/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceV2ResponseData
type QueryInvoiceV2ResponseData struct {
	//
	InvoiceInfoList []*QueryInvoiceV2ResponseDataInvoiceInfoListInner `json:"invoice_info_list"`
	PageInfo        *QueryInvoiceV2ResponseDataPageInfo               `json:"page_info,omitempty"`
}
