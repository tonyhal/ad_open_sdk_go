/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeGetV2ResponseDataListInner struct for CreativeGetV2ResponseDataListInner
type CreativeGetV2ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	CreativeCreateTime *string `json:"creative_create_time,omitempty"`
	//
	CreativeId *int64 `json:"creative_id,omitempty"`
	//
	CreativeModifyTime *string `json:"creative_modify_time,omitempty"`
	//
	CreativeWordIds []int64 `json:"creative_word_ids,omitempty"`
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	ImageIds  []string                        `json:"image_ids,omitempty"`
	ImageMode *CreativeGetV2DataListImageMode `json:"image_mode,omitempty"`
	//
	Materials []*CreativeGetV2ResponseDataListInnerMaterialsInner `json:"materials,omitempty"`
	OptStatus *CreativeGetV2DataListOptStatus                     `json:"opt_status,omitempty"`
	Status    *CreativeGetV2DataListStatus                        `json:"status,omitempty"`
	//
	ThirdPartyId *string `json:"third_party_id,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
