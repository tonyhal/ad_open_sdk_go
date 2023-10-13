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

// AudiencePackageUpdateV2HideIfExists
type AudiencePackageUpdateV2HideIfExists int64

// List of audience_package_update_v2_hide_if_exists
const (
	Enum_0_AudiencePackageUpdateV2HideIfExists AudiencePackageUpdateV2HideIfExists = 0
	Enum_1_AudiencePackageUpdateV2HideIfExists AudiencePackageUpdateV2HideIfExists = 1
	Enum_2_AudiencePackageUpdateV2HideIfExists AudiencePackageUpdateV2HideIfExists = 2
)

// All allowed values of AudiencePackageUpdateV2HideIfExists enum
var AllowedAudiencePackageUpdateV2HideIfExistsEnumValues = []AudiencePackageUpdateV2HideIfExists{
	0,
	1,
	2,
}

// NewAudiencePackageUpdateV2HideIfExistsFromValue returns a pointer to a valid AudiencePackageUpdateV2HideIfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2HideIfExistsFromValue(v int64) (*AudiencePackageUpdateV2HideIfExists, error) {
	ev := AudiencePackageUpdateV2HideIfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2HideIfExists: valid values are %v", v, AllowedAudiencePackageUpdateV2HideIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2HideIfExists) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2HideIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_hide_if_exists value
func (v AudiencePackageUpdateV2HideIfExists) Ptr() *AudiencePackageUpdateV2HideIfExists {
	return &v
}
