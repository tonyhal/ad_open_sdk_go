/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolQuickAppManagementQuickAppGetV2ResponseDataQuickAppInfoInner struct for ToolQuickAppManagementQuickAppGetV2ResponseDataQuickAppInfoInner
type ToolQuickAppManagementQuickAppGetV2ResponseDataQuickAppInfoInner struct {
	//
	CreateTime string `json:"create_time"`
	//
	HomepageUrl string `json:"homepage_url"`
	//
	Name string `json:"name"`
	//
	PackageName string `json:"package_name"`
	//
	QuickAppId int64                                                     `json:"quick_app_id"`
	Status     ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus `json:"status"`
	//
	UpdateTime string `json:"update_time"`
}
