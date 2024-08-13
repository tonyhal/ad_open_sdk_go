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

// AdGetV2DataConvertedTimeDuration
type AdGetV2DataConvertedTimeDuration string

// List of ad_get_v2_data_converted_time_duration
const (
	THREE_MONTH_AdGetV2DataConvertedTimeDuration  AdGetV2DataConvertedTimeDuration = "THREE_MONTH"
	SEVEN_DAY_AdGetV2DataConvertedTimeDuration    AdGetV2DataConvertedTimeDuration = "SEVEN_DAY"
	TWELVE_MONTH_AdGetV2DataConvertedTimeDuration AdGetV2DataConvertedTimeDuration = "TWELVE_MONTH"
	TODAY_AdGetV2DataConvertedTimeDuration        AdGetV2DataConvertedTimeDuration = "TODAY"
	ONE_MONTH_AdGetV2DataConvertedTimeDuration    AdGetV2DataConvertedTimeDuration = "ONE_MONTH"
	SIX_MONTH_AdGetV2DataConvertedTimeDuration    AdGetV2DataConvertedTimeDuration = "SIX_MONTH"
	NONE_AdGetV2DataConvertedTimeDuration         AdGetV2DataConvertedTimeDuration = "NONE"
)

// All allowed values of AdGetV2DataConvertedTimeDuration enum
var AllowedAdGetV2DataConvertedTimeDurationEnumValues = []AdGetV2DataConvertedTimeDuration{
	"THREE_MONTH",
	"SEVEN_DAY",
	"TWELVE_MONTH",
	"TODAY",
	"ONE_MONTH",
	"SIX_MONTH",
	"NONE",
}

// NewAdGetV2DataConvertedTimeDurationFromValue returns a pointer to a valid AdGetV2DataConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataConvertedTimeDurationFromValue(v string) (*AdGetV2DataConvertedTimeDuration, error) {
	ev := AdGetV2DataConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataConvertedTimeDuration: valid values are %v", v, AllowedAdGetV2DataConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_converted_time_duration value
func (v AdGetV2DataConvertedTimeDuration) Ptr() *AdGetV2DataConvertedTimeDuration {
	return &v
}
