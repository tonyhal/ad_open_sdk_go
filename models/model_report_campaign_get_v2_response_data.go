/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2ResponseData
type ReportCampaignGetV2ResponseData struct {
	//
	List     []*ReportCampaignGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportCampaignGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
