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

// AdConvertOptimizedTargetGetV30ConvertTracking
type AdConvertOptimizedTargetGetV30ConvertTracking string

// List of ad_convert_optimized_target_get_v3.0_convert_tracking
const (
	AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD_AdConvertOptimizedTargetGetV30ConvertTracking AdConvertOptimizedTargetGetV30ConvertTracking = "AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD"
	AD_CONVERT_SOURCE_TYPE_OPEN_URL_AdConvertOptimizedTargetGetV30ConvertTracking     AdConvertOptimizedTargetGetV30ConvertTracking = "AD_CONVERT_SOURCE_TYPE_OPEN_URL"
	AD_CONVERT_SOURCE_TYPE_SDK_AdConvertOptimizedTargetGetV30ConvertTracking          AdConvertOptimizedTargetGetV30ConvertTracking = "AD_CONVERT_SOURCE_TYPE_SDK"
)

// All allowed values of AdConvertOptimizedTargetGetV30ConvertTracking enum
var AllowedAdConvertOptimizedTargetGetV30ConvertTrackingEnumValues = []AdConvertOptimizedTargetGetV30ConvertTracking{
	"AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD",
	"AD_CONVERT_SOURCE_TYPE_OPEN_URL",
	"AD_CONVERT_SOURCE_TYPE_SDK",
}

// NewAdConvertOptimizedTargetGetV30ConvertTrackingFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30ConvertTracking
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30ConvertTrackingFromValue(v string) (*AdConvertOptimizedTargetGetV30ConvertTracking, error) {
	ev := AdConvertOptimizedTargetGetV30ConvertTracking(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30ConvertTracking: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30ConvertTrackingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30ConvertTracking) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30ConvertTrackingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_convert_tracking value
func (v AdConvertOptimizedTargetGetV30ConvertTracking) Ptr() *AdConvertOptimizedTargetGetV30ConvertTracking {
	return &v
}
