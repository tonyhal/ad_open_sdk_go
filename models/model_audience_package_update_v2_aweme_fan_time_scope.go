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

// AudiencePackageUpdateV2AwemeFanTimeScope
type AudiencePackageUpdateV2AwemeFanTimeScope string

// List of audience_package_update_v2_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_AudiencePackageUpdateV2AwemeFanTimeScope AudiencePackageUpdateV2AwemeFanTimeScope = "FIFTEEN_DAYS"
	SIXTY_DAYS_AudiencePackageUpdateV2AwemeFanTimeScope   AudiencePackageUpdateV2AwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_AudiencePackageUpdateV2AwemeFanTimeScope  AudiencePackageUpdateV2AwemeFanTimeScope = "THIRTY_DAYS"
)

// All allowed values of AudiencePackageUpdateV2AwemeFanTimeScope enum
var AllowedAudiencePackageUpdateV2AwemeFanTimeScopeEnumValues = []AudiencePackageUpdateV2AwemeFanTimeScope{
	"FIFTEEN_DAYS",
	"SIXTY_DAYS",
	"THIRTY_DAYS",
}

// NewAudiencePackageUpdateV2AwemeFanTimeScopeFromValue returns a pointer to a valid AudiencePackageUpdateV2AwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2AwemeFanTimeScopeFromValue(v string) (*AudiencePackageUpdateV2AwemeFanTimeScope, error) {
	ev := AudiencePackageUpdateV2AwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2AwemeFanTimeScope: valid values are %v", v, AllowedAudiencePackageUpdateV2AwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2AwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2AwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_aweme_fan_time_scope value
func (v AudiencePackageUpdateV2AwemeFanTimeScope) Ptr() *AudiencePackageUpdateV2AwemeFanTimeScope {
	return &v
}
