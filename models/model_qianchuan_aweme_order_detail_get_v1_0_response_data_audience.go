/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataAudience 人群定向
type QianchuanAwemeOrderDetailGetV10ResponseDataAudience struct {
	// 年龄定向
	Age          []*QianchuanAwemeOrderDetailGetV10DataAudienceAge        `json:"age,omitempty"`
	AudienceMode *QianchuanAwemeOrderDetailGetV10DataAudienceAudienceMode `json:"audience_mode,omitempty"`
	// 抖音达人ID列表
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	// 抖音用户行为类型 ，
	AwemeFanBehaviors []*QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	// 具体定向的城市列表
	City                 []int64                                                          `json:"city,omitempty"`
	District             *QianchuanAwemeOrderDetailGetV10DataAudienceDistrict             `json:"district,omitempty"`
	ExcludeLimitedRegion *QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAwemeOrderDetailGetV10DataAudienceGender               `json:"gender,omitempty"`
	// 兴趣类目词
	InterestCategories []int64 `json:"interest_categories,omitempty"`
}
