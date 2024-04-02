/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30Filtering
type AdlabGroupListV30Filtering struct {
	// 按广告组ID进行过滤
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	// 按照项目名称进行过滤
	GroupName *string `json:"group_name,omitempty"`
	// 按项目ID进行过滤
	Ids         []int64                                `json:"ids,omitempty"`
	LandingType *AdlabGroupListV30FilteringLandingType `json:"landing_type,omitempty"`
}
