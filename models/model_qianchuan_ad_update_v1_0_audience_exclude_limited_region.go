/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceExcludeLimitedRegion
type QianchuanAdUpdateV10AudienceExcludeLimitedRegion int64

// List of qianchuan_ad_update_v1.0_audience_exclude_limited_region
const (
	Enum_0_QianchuanAdUpdateV10AudienceExcludeLimitedRegion QianchuanAdUpdateV10AudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanAdUpdateV10AudienceExcludeLimitedRegion QianchuanAdUpdateV10AudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAdUpdateV10AudienceExcludeLimitedRegion enum
var AllowedQianchuanAdUpdateV10AudienceExcludeLimitedRegionEnumValues = []QianchuanAdUpdateV10AudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAdUpdateV10AudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanAdUpdateV10AudienceExcludeLimitedRegion, error) {
	ev := QianchuanAdUpdateV10AudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_exclude_limited_region value
func (v QianchuanAdUpdateV10AudienceExcludeLimitedRegion) Ptr() *QianchuanAdUpdateV10AudienceExcludeLimitedRegion {
	return &v
}
