/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMaterialPushV30Request struct for AicMaterialPushV30Request
type AicMaterialPushV30Request struct {
	// BP账户体系:组织id，Agent账户体系:代理商id
	AccountId   int64                         `json:"account_id"`
	AccountType AicMaterialPushV30AccountType `json:"account_type"`
	// 待推送的广告主，数量限制：50，必须是推送广告主所属平台的广告主
	AdvertiserIds []int64                         `json:"advertiser_ids"`
	BussinessType AicMaterialPushV30BussinessType `json:"bussiness_type"`
	// 需要推送的视频，数量限制50
	Videos []*AicMaterialPushV30RequestVideosInner `json:"videos"`
}