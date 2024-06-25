/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2LocationType
type AudiencePackageCreateV2LocationType string

// List of audience_package_create_v2_location_type
const (
	TRAVEL_AudiencePackageCreateV2LocationType  AudiencePackageCreateV2LocationType = "TRAVEL"
	HOME_AudiencePackageCreateV2LocationType    AudiencePackageCreateV2LocationType = "HOME"
	CURRENT_AudiencePackageCreateV2LocationType AudiencePackageCreateV2LocationType = "CURRENT"
	ALL_AudiencePackageCreateV2LocationType     AudiencePackageCreateV2LocationType = "ALL"
)

// All allowed values of AudiencePackageCreateV2LocationType enum
var AllowedAudiencePackageCreateV2LocationTypeEnumValues = []AudiencePackageCreateV2LocationType{
	"TRAVEL",
	"HOME",
	"CURRENT",
	"ALL",
}

// NewAudiencePackageCreateV2LocationTypeFromValue returns a pointer to a valid AudiencePackageCreateV2LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2LocationTypeFromValue(v string) (*AudiencePackageCreateV2LocationType, error) {
	ev := AudiencePackageCreateV2LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2LocationType: valid values are %v", v, AllowedAudiencePackageCreateV2LocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2LocationType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2LocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_location_type value
func (v AudiencePackageCreateV2LocationType) Ptr() *AudiencePackageCreateV2LocationType {
	return &v
}
