/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackageReportV2Filter
type ToolsUnionFlowPackageReportV2Filter struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	EndTime **string `json:"end_time,omitempty"`
	//
	Rits []int64 `json:"rits,omitempty"`
	//
	StartTime **string `json:"start_time,omitempty"`
}
