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

// AdGetV2DataAudienceDeviceType
type AdGetV2DataAudienceDeviceType string

// List of ad_get_v2_data_audience_device_type
const (
	PAD_AdGetV2DataAudienceDeviceType    AdGetV2DataAudienceDeviceType = "PAD"
	MOBILE_AdGetV2DataAudienceDeviceType AdGetV2DataAudienceDeviceType = "MOBILE"
)

// All allowed values of AdGetV2DataAudienceDeviceType enum
var AllowedAdGetV2DataAudienceDeviceTypeEnumValues = []AdGetV2DataAudienceDeviceType{
	"PAD",
	"MOBILE",
}

// NewAdGetV2DataAudienceDeviceTypeFromValue returns a pointer to a valid AdGetV2DataAudienceDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDeviceTypeFromValue(v string) (*AdGetV2DataAudienceDeviceType, error) {
	ev := AdGetV2DataAudienceDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDeviceType: valid values are %v", v, AllowedAdGetV2DataAudienceDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDeviceType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_device_type value
func (v AdGetV2DataAudienceDeviceType) Ptr() *AdGetV2DataAudienceDeviceType {
	return &v
}
