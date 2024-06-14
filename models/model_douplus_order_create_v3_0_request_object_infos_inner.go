/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderCreateV30RequestObjectInfosInner struct for DouplusOrderCreateV30RequestObjectInfosInner
type DouplusOrderCreateV30RequestObjectInfosInner struct {
	// 抖音号，即客户在手机端上看到的抖音号，若向客户披露抖音号请使用该字段
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音视频ID，当营销目标为视频加热和视频加热直播间：传递视频ID
	ItemId        *int64                                        `json:"item_id,omitempty"`
	MarketingGoal DouplusOrderCreateV30ObjectInfosMarketingGoal `json:"marketing_goal"`
}
