/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceLocationType
type AdGetV2DataAudienceLocationType string

// List of ad_get_v2_data_audience_location_type
const (
	HOME_AdGetV2DataAudienceLocationType    AdGetV2DataAudienceLocationType = "HOME"
	CURRENT_AdGetV2DataAudienceLocationType AdGetV2DataAudienceLocationType = "CURRENT"
	ALL_AdGetV2DataAudienceLocationType     AdGetV2DataAudienceLocationType = "ALL"
	TRAVEL_AdGetV2DataAudienceLocationType  AdGetV2DataAudienceLocationType = "TRAVEL"
)

// All allowed values of AdGetV2DataAudienceLocationType enum
var AllowedAdGetV2DataAudienceLocationTypeEnumValues = []AdGetV2DataAudienceLocationType{
	"HOME",
	"CURRENT",
	"ALL",
	"TRAVEL",
}

// NewAdGetV2DataAudienceLocationTypeFromValue returns a pointer to a valid AdGetV2DataAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceLocationTypeFromValue(v string) (*AdGetV2DataAudienceLocationType, error) {
	ev := AdGetV2DataAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceLocationType: valid values are %v", v, AllowedAdGetV2DataAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_location_type value
func (v AdGetV2DataAudienceLocationType) Ptr() *AdGetV2DataAudienceLocationType {
	return &v
}
