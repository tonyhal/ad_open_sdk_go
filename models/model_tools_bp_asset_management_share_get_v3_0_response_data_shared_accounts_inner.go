/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareGetV30ResponseDataSharedAccountsInner struct for ToolsBpAssetManagementShareGetV30ResponseDataSharedAccountsInner
type ToolsBpAssetManagementShareGetV30ResponseDataSharedAccountsInner struct {
	AccountInfo          *ToolsBpAssetManagementShareGetV30ResponseDataSharedAccountsInnerAccountInfo          `json:"account_info,omitempty"`
	AllAccountsByBp      *ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp                   `json:"all_accounts_by_bp,omitempty"`
	AllAccountsByCompany *ToolsBpAssetManagementShareGetV30ResponseDataSharedAccountsInnerAllAccountsByCompany `json:"all_accounts_by_company,omitempty"`
	ShareMode            ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode                          `json:"share_mode"`
}
