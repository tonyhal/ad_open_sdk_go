/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInner struct for PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInner
type PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	Creatives []*PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInner `json:"creatives,omitempty"`
}
