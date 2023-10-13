/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataUnionVideoType
type AdGetV2DataUnionVideoType string

// List of ad_get_v2_data_union_video_type
const (
	SPLASH_VIDEO_AdGetV2DataUnionVideoType   AdGetV2DataUnionVideoType = "SPLASH_VIDEO"
	REWARDED_VIDEO_AdGetV2DataUnionVideoType AdGetV2DataUnionVideoType = "REWARDED_VIDEO"
	ORIGINAL_VIDEO_AdGetV2DataUnionVideoType AdGetV2DataUnionVideoType = "ORIGINAL_VIDEO"
)

// All allowed values of AdGetV2DataUnionVideoType enum
var AllowedAdGetV2DataUnionVideoTypeEnumValues = []AdGetV2DataUnionVideoType{
	"SPLASH_VIDEO",
	"REWARDED_VIDEO",
	"ORIGINAL_VIDEO",
}

// NewAdGetV2DataUnionVideoTypeFromValue returns a pointer to a valid AdGetV2DataUnionVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataUnionVideoTypeFromValue(v string) (*AdGetV2DataUnionVideoType, error) {
	ev := AdGetV2DataUnionVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataUnionVideoType: valid values are %v", v, AllowedAdGetV2DataUnionVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataUnionVideoType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataUnionVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_union_video_type value
func (v AdGetV2DataUnionVideoType) Ptr() *AdGetV2DataUnionVideoType {
	return &v
}
