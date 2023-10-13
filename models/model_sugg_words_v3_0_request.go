/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SuggWordsV30Request struct for SuggWordsV30Request
type SuggWordsV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目id
	ProjectId          *int64                                 `json:"project_id,omitempty"`
	PromotionMaterials *SuggWordsV30RequestPromotionMaterials `json:"promotion_materials,omitempty"`
	// 以词推词，最多只传入10个，长度要求30字内 ,可传入行业名称获取行业推词
	QueryList []string `json:"query_list,omitempty"`
}
