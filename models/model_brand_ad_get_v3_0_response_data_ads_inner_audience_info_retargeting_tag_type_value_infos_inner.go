/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerAudienceInfoRetargetingTagTypeValueInfosInner struct for BrandAdGetV30ResponseDataAdsInnerAudienceInfoRetargetingTagTypeValueInfosInner
type BrandAdGetV30ResponseDataAdsInnerAudienceInfoRetargetingTagTypeValueInfosInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AppIdList []int64 `json:"app_id_list,omitempty"`
	//
	CategoryList []int64 `json:"category_list,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	ExpireTime int64 `json:"expire_time"`
	//
	Id int64 `json:"id"`
	//
	Name string `json:"name"`
	//
	PullOffTagStatus *int64 `json:"pull_off_tag_status,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
