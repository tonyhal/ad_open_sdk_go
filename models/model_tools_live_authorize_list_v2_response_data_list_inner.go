/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLiveAuthorizeListV2ResponseDataListInner struct for ToolsLiveAuthorizeListV2ResponseDataListInner
type ToolsLiveAuthorizeListV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AuthorizeEndTime *string `json:"authorize_end_time,omitempty"`
	//
	AuthorizeStartTime *string `json:"authorize_start_time,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	IesId *string `json:"ies_id,omitempty"`
	//
	IesUserName *string `json:"ies_user_name,omitempty"`
	//
	LimitedPromotionTypes []*ToolsLiveAuthorizeListV2ResponseDataListInnerLimitedPromotionTypesInner `json:"limited_promotion_types,omitempty"`
	Status                *ToolsLiveAuthorizeListV2DataListStatus                                    `json:"status,omitempty"`
}
