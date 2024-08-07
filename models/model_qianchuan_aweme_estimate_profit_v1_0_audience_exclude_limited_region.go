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

// QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion
type QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion int64

// List of qianchuan_aweme_estimate_profit_v1.0_audience_exclude_limited_region
const (
	Enum_0_QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion enum
var AllowedQianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegionEnumValues = []QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion, error) {
	ev := QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_audience_exclude_limited_region value
func (v QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion) Ptr() *QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion {
	return &v
}
