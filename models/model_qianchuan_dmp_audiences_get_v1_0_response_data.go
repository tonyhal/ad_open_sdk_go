/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanDmpAudiencesGetV10ResponseData
type QianchuanDmpAudiencesGetV10ResponseData struct {
	//
	Offset *int64 `json:"offset,omitempty"`
	//
	RetargetingTags []*QianchuanDmpAudiencesGetV10ResponseDataRetargetingTagsInner `json:"retargeting_tags,omitempty"`
	//
	TotalNum *int64 `json:"total_num,omitempty"`
}
