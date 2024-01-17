/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdConvertOptimizedTargetGetV30MarketingPurpose
type AdConvertOptimizedTargetGetV30MarketingPurpose string

// List of ad_convert_optimized_target_get_v3.0_marketing_purpose
const (
	CONVERSION_AdConvertOptimizedTargetGetV30MarketingPurpose AdConvertOptimizedTargetGetV30MarketingPurpose = "CONVERSION"
)

// All allowed values of AdConvertOptimizedTargetGetV30MarketingPurpose enum
var AllowedAdConvertOptimizedTargetGetV30MarketingPurposeEnumValues = []AdConvertOptimizedTargetGetV30MarketingPurpose{
	"CONVERSION",
}

// NewAdConvertOptimizedTargetGetV30MarketingPurposeFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30MarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30MarketingPurposeFromValue(v string) (*AdConvertOptimizedTargetGetV30MarketingPurpose, error) {
	ev := AdConvertOptimizedTargetGetV30MarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30MarketingPurpose: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30MarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30MarketingPurpose) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30MarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_marketing_purpose value
func (v AdConvertOptimizedTargetGetV30MarketingPurpose) Ptr() *AdConvertOptimizedTargetGetV30MarketingPurpose {
	return &v
}