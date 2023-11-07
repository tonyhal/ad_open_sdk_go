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

// AdGetV2DataAudienceAwemeFanTimeScope
type AdGetV2DataAudienceAwemeFanTimeScope string

// List of ad_get_v2_data_audience_aweme_fan_time_scope
const (
	THIRTY_DAYS_AdGetV2DataAudienceAwemeFanTimeScope  AdGetV2DataAudienceAwemeFanTimeScope = "THIRTY_DAYS"
	SIXTY_DAYS_AdGetV2DataAudienceAwemeFanTimeScope   AdGetV2DataAudienceAwemeFanTimeScope = "SIXTY_DAYS"
	FIFTEEN_DAYS_AdGetV2DataAudienceAwemeFanTimeScope AdGetV2DataAudienceAwemeFanTimeScope = "FIFTEEN_DAYS"
)

// All allowed values of AdGetV2DataAudienceAwemeFanTimeScope enum
var AllowedAdGetV2DataAudienceAwemeFanTimeScopeEnumValues = []AdGetV2DataAudienceAwemeFanTimeScope{
	"THIRTY_DAYS",
	"SIXTY_DAYS",
	"FIFTEEN_DAYS",
}

// NewAdGetV2DataAudienceAwemeFanTimeScopeFromValue returns a pointer to a valid AdGetV2DataAudienceAwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAwemeFanTimeScopeFromValue(v string) (*AdGetV2DataAudienceAwemeFanTimeScope, error) {
	ev := AdGetV2DataAudienceAwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAwemeFanTimeScope: valid values are %v", v, AllowedAdGetV2DataAudienceAwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_aweme_fan_time_scope value
func (v AdGetV2DataAudienceAwemeFanTimeScope) Ptr() *AdGetV2DataAudienceAwemeFanTimeScope {
	return &v
}
