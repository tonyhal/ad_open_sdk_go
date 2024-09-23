/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectCreateV30RequestAudience
type LocalProjectCreateV30RequestAudience struct {
	// 年龄
	Age                   []*LocalProjectCreateV30AudienceAge                 `json:"age,omitempty"`
	ConvertedTimeDuration *LocalProjectCreateV30AudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	CustomArea            *LocalProjectCreateV30RequestAudienceCustomArea     `json:"custom_area,omitempty"`
	District              LocalProjectCreateV30AudienceDistrict               `json:"district"`
	Gender                *LocalProjectCreateV30AudienceGender                `json:"gender,omitempty"`
	HideIfConverted       *LocalProjectCreateV30AudienceHideIfConverted       `json:"hide_if_converted,omitempty"`
	PoiAround             *LocalProjectCreateV30RequestAudiencePoiAround      `json:"poi_around,omitempty"`
	Region                *LocalProjectCreateV30RequestAudienceRegion         `json:"region,omitempty"`
	// 定向人群包ID
	RetargetingTags []int64 `json:"retargeting_tags,omitempty"`
	// 排除人群包ID
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
}
