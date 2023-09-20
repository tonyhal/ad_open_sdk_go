/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2Ac
type AudiencePackageCreateV2Ac string

// List of audience_package_create_v2_ac
const (
	Enum_4_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "4G"
	Enum_2_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "2G"
	WIFI_AudiencePackageCreateV2Ac     AudiencePackageCreateV2Ac = "WIFI"
	Enum_3_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "3G"
)

// All allowed values of AudiencePackageCreateV2Ac enum
var AllowedAudiencePackageCreateV2AcEnumValues = []AudiencePackageCreateV2Ac{
	"4G",
	"2G",
	"WIFI",
	"3G",
}

// NewAudiencePackageCreateV2AcFromValue returns a pointer to a valid AudiencePackageCreateV2Ac
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AcFromValue(v string) (*AudiencePackageCreateV2Ac, error) {
	ev := AudiencePackageCreateV2Ac(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2Ac: valid values are %v", v, AllowedAudiencePackageCreateV2AcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2Ac) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_ac value
func (v AudiencePackageCreateV2Ac) Ptr() *AudiencePackageCreateV2Ac {
	return &v
}
