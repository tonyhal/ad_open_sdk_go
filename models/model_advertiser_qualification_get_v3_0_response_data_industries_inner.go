/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationGetV30ResponseDataIndustriesInner struct for AdvertiserQualificationGetV30ResponseDataIndustriesInner
type AdvertiserQualificationGetV30ResponseDataIndustriesInner struct {
	// 一级行业ID
	Industry1Id *int64 `json:"industry1_id,omitempty"`
	// 一级行业名称
	Industry1Name *string `json:"industry1_name,omitempty"`
	// 二级行业ID
	Industry2Id *int64 `json:"industry2_id,omitempty"`
	// 二级行业名称
	Industry2Name *string `json:"industry2_name,omitempty"`
	// 三级行业ID
	Industry3Id *int64 `json:"industry3_id,omitempty"`
	// 三级行业名称
	Industry3Name *string `json:"industry3_name,omitempty"`
	// 是否是历史的补充资质及推广资质，如果是，则行业相关字段无值，重新提交时需归档到具体行业下
	IsHistory *bool `json:"is_history,omitempty"`
	// 开户资质列表，相关字段见下
	Others    []*AdvertiserQualificationGetV30ResponseDataIndustriesInnerOthersInner `json:"others,omitempty"`
	Promotion *AdvertiserQualificationGetV30ResponseDataIndustriesInnerPromotion     `json:"promotion,omitempty"`
}
