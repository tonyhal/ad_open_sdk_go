/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeItemsDataV2ResponseData
type StarDemandOmGetChallengeItemsDataV2ResponseData struct {
	//
	DataList []*StarDemandOmGetChallengeItemsDataV2ResponseDataDataListInner `json:"data_list,omitempty"`
	PageInfo *StarDemandOmGetChallengeItemsDataV2ResponseDataPageInfo        `json:"page_info,omitempty"`
}