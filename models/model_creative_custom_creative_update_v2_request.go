/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2Request struct for CreativeCustomCreativeUpdateV2Request
type CreativeCustomCreativeUpdateV2Request struct {
	AdData CreativeCustomCreativeUpdateV2RequestAdData `json:"ad_data"`
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CreativeList []*CreativeCustomCreativeUpdateV2RequestCreativeListInner `json:"creative_list"`
}
