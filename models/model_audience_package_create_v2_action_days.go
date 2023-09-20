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

// AudiencePackageCreateV2ActionDays
type AudiencePackageCreateV2ActionDays int64

// List of audience_package_create_v2_action_days
const (
	Enum_7_AudiencePackageCreateV2ActionDays   AudiencePackageCreateV2ActionDays = 7
	Enum_15_AudiencePackageCreateV2ActionDays  AudiencePackageCreateV2ActionDays = 15
	Enum_30_AudiencePackageCreateV2ActionDays  AudiencePackageCreateV2ActionDays = 30
	Enum_60_AudiencePackageCreateV2ActionDays  AudiencePackageCreateV2ActionDays = 60
	Enum_90_AudiencePackageCreateV2ActionDays  AudiencePackageCreateV2ActionDays = 90
	Enum_180_AudiencePackageCreateV2ActionDays AudiencePackageCreateV2ActionDays = 180
	Enum_365_AudiencePackageCreateV2ActionDays AudiencePackageCreateV2ActionDays = 365
)

// All allowed values of AudiencePackageCreateV2ActionDays enum
var AllowedAudiencePackageCreateV2ActionDaysEnumValues = []AudiencePackageCreateV2ActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewAudiencePackageCreateV2ActionDaysFromValue returns a pointer to a valid AudiencePackageCreateV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2ActionDaysFromValue(v int64) (*AudiencePackageCreateV2ActionDays, error) {
	ev := AudiencePackageCreateV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2ActionDays: valid values are %v", v, AllowedAudiencePackageCreateV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2ActionDays) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_action_days value
func (v AudiencePackageCreateV2ActionDays) Ptr() *AudiencePackageCreateV2ActionDays {
	return &v
}
