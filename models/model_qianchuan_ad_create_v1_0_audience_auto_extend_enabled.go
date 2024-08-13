/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceAutoExtendEnabled
type QianchuanAdCreateV10AudienceAutoExtendEnabled int64

// List of qianchuan_ad_create_v1.0_audience_auto_extend_enabled
const (
	Enum_0_QianchuanAdCreateV10AudienceAutoExtendEnabled QianchuanAdCreateV10AudienceAutoExtendEnabled = 0
	Enum_1_QianchuanAdCreateV10AudienceAutoExtendEnabled QianchuanAdCreateV10AudienceAutoExtendEnabled = 1
)

// All allowed values of QianchuanAdCreateV10AudienceAutoExtendEnabled enum
var AllowedQianchuanAdCreateV10AudienceAutoExtendEnabledEnumValues = []QianchuanAdCreateV10AudienceAutoExtendEnabled{
	0,
	1,
}

// NewQianchuanAdCreateV10AudienceAutoExtendEnabledFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAutoExtendEnabled
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAutoExtendEnabledFromValue(v int64) (*QianchuanAdCreateV10AudienceAutoExtendEnabled, error) {
	ev := QianchuanAdCreateV10AudienceAutoExtendEnabled(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAutoExtendEnabled: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAutoExtendEnabledEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAutoExtendEnabled) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAutoExtendEnabledEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_auto_extend_enabled value
func (v QianchuanAdCreateV10AudienceAutoExtendEnabled) Ptr() *QianchuanAdCreateV10AudienceAutoExtendEnabled {
	return &v
}
