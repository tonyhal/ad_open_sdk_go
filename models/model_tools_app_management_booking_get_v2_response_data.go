/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBookingGetV2ResponseData
type ToolsAppManagementBookingGetV2ResponseData struct {
	//
	List     []*ToolsAppManagementBookingGetV2ResponseDataListInner `json:"list"`
	PageInfo *ToolsAppManagementBookingGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
