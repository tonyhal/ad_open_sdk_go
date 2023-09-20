/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2ActionDays
type AudiencePackageUpdateV2ActionDays int64

// List of audience_package_update_v2_action_days
const (
	Enum_7_AudiencePackageUpdateV2ActionDays   AudiencePackageUpdateV2ActionDays = 7
	Enum_15_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 15
	Enum_30_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 30
	Enum_60_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 60
	Enum_90_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 90
	Enum_180_AudiencePackageUpdateV2ActionDays AudiencePackageUpdateV2ActionDays = 180
	Enum_365_AudiencePackageUpdateV2ActionDays AudiencePackageUpdateV2ActionDays = 365
)

// All allowed values of AudiencePackageUpdateV2ActionDays enum
var AllowedAudiencePackageUpdateV2ActionDaysEnumValues = []AudiencePackageUpdateV2ActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewAudiencePackageUpdateV2ActionDaysFromValue returns a pointer to a valid AudiencePackageUpdateV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2ActionDaysFromValue(v int64) (*AudiencePackageUpdateV2ActionDays, error) {
	ev := AudiencePackageUpdateV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2ActionDays: valid values are %v", v, AllowedAudiencePackageUpdateV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2ActionDays) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_action_days value
func (v AudiencePackageUpdateV2ActionDays) Ptr() *AudiencePackageUpdateV2ActionDays {
	return &v
}
