/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdCreateV10Request struct for QianchuanUniAwemeAdCreateV10Request
type QianchuanUniAwemeAdCreateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 抖音号id
	AwemeId                       int64                                                             `json:"aweme_id"`
	CreativeSetting               *QianchuanUniAwemeAdCreateV10RequestCreativeSetting               `json:"creative_setting,omitempty"`
	DeliverySetting               QianchuanUniAwemeAdCreateV10RequestDeliverySetting                `json:"delivery_setting"`
	MarketingGoal                 QianchuanUniAwemeAdCreateV10MarketingGoal                         `json:"marketing_goal"`
	ProgrammaticCreativeMediaList *QianchuanUniAwemeAdCreateV10RequestProgrammaticCreativeMediaList `json:"programmatic_creative_media_list,omitempty"`
}
