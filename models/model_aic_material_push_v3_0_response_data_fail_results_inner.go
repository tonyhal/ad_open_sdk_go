/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMaterialPushV30ResponseDataFailResultsInner struct for AicMaterialPushV30ResponseDataFailResultsInner
type AicMaterialPushV30ResponseDataFailResultsInner struct {
	// 推送失败的原因
	Reason *string `json:"reason,omitempty"`
	// 目标推送广告主id
	TargetAdvertiserId *int64 `json:"target_advertiser_id,omitempty"`
	// 推送失败的视频id
	VideoId *string `json:"video_id,omitempty"`
}