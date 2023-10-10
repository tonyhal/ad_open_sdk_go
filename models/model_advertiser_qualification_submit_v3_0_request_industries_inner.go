/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationSubmitV30RequestIndustriesInner struct for AdvertiserQualificationSubmitV30RequestIndustriesInner
type AdvertiserQualificationSubmitV30RequestIndustriesInner struct {
	// 行业id
	IndustryId int64 `json:"industry_id"`
	// 行业补充资质列表，相关字段见下
	Others    []*AdvertiserQualificationSubmitV30RequestIndustriesInnerOthersInner `json:"others,omitempty"`
	Promotion AdvertiserQualificationSubmitV30RequestIndustriesInnerPromotion      `json:"promotion"`
}
