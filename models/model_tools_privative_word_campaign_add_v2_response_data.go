/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordCampaignAddV2ResponseData
type ToolsPrivativeWordCampaignAddV2ResponseData struct {
	//
	CampaignErrorList []int64 `json:"campaign_error_list,omitempty"`
	//
	CampaignList []*ToolsPrivativeWordCampaignAddV2ResponseDataCampaignListInner `json:"campaign_list,omitempty"`
}
