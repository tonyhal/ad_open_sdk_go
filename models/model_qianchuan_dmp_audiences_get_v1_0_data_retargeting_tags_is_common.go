/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon
type QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon int64

// List of qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_is_common
const (
	Enum_0_QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon = 0
	Enum_1_QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon = 1
)

// All allowed values of QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon enum
var AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommonEnumValues = []QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon{
	0,
	1,
}

// NewQianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommonFromValue returns a pointer to a valid QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommonFromValue(v int64) (*QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon, error) {
	ev := QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon: valid values are %v", v, AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon) IsValid() bool {
	for _, existing := range AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_is_common value
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon) Ptr() *QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon {
	return &v
}
