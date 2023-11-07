/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType
type QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType string

// List of qianchuan_ad_detail_get_v1.0_data_audience_inactive_retargeting_tags_inactive_type
const (
	EXPIRE_QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType         QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType = "EXPIRE"
	MANUAL_OFFLINE_QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType = "MANUAL_OFFLINE"
	TAG_OFFLINE_QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType    QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType = "TAG_OFFLINE"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType enum
var AllowedQianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveTypeEnumValues = []QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType{
	"EXPIRE",
	"MANUAL_OFFLINE",
	"TAG_OFFLINE",
}

// NewQianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveTypeFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType, error) {
	ev := QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_inactive_retargeting_tags_inactive_type value
func (v QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType) Ptr() *QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType {
	return &v
}
