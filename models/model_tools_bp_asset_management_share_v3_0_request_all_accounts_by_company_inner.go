/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareV30RequestAllAccountsByCompanyInner struct for ToolsBpAssetManagementShareV30RequestAllAccountsByCompanyInner
type ToolsBpAssetManagementShareV30RequestAllAccountsByCompanyInner struct {
	AccountType ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType `json:"account_type"`
	// 公司ID
	CompanyId int64 `json:"company_id"`
}