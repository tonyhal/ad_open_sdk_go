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

// QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion
type QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion int64

// List of qianchuan_aweme_order_detail_get_v1.0_data_audience_exclude_limited_region
const (
	Enum_0_QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion enum
var AllowedQianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegionEnumValues = []QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_audience_exclude_limited_region value
func (v QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion) Ptr() *QianchuanAwemeOrderDetailGetV10DataAudienceExcludeLimitedRegion {
	return &v
}
