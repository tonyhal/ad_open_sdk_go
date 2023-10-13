/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBpShareV2Request struct for ToolsAppManagementBpShareV2Request
type ToolsAppManagementBpShareV2Request struct {
	//
	AccountInfos        []*ToolsAppManagementBpShareV2RequestAccountInfosInner `json:"account_infos,omitempty"`
	AllAccountByCompany *ToolsAppManagementBpShareV2RequestAllAccountByCompany `json:"all_account_by_company,omitempty"`
	//
	AllAccounts []*ToolsAppManagementBpShareV2RequestAllAccountsInner `json:"all_accounts,omitempty"`
	//
	OrganizationId int64 `json:"organization_id"`
	//
	PackageId string                               `json:"package_id"`
	ShareMode ToolsAppManagementBpShareV2ShareMode `json:"share_mode"`
}
