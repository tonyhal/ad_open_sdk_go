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

// QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion
type QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion int64

// List of qianchuan_aweme_order_create_v1.0_audience_exclude_limited_region
const (
	Enum_0_QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion enum
var AllowedQianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegionEnumValues = []QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion, error) {
	ev := QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_audience_exclude_limited_region value
func (v QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion) Ptr() *QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion {
	return &v
}
