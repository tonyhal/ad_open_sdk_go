/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreateStatementInvoiceV2RequestInvoiceBillListInner struct for CreateStatementInvoiceV2RequestInvoiceBillListInner
type CreateStatementInvoiceV2RequestInvoiceBillListInner struct {
	// 开票票据项目列表
	InvoiceBillProjectList []*CreateStatementInvoiceV2RequestInvoiceBillListInnerInvoiceBillProjectListInner `json:"invoice_bill_project_list"`
	// 打印备注
	Remark *string `json:"remark,omitempty"`
}
