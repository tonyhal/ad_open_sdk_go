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

// AudiencePackageGetV2DataAudiencePackagesAudienceLocationType
type AudiencePackageGetV2DataAudiencePackagesAudienceLocationType string

// List of audience_package_get_v2_data_audience_packages_audience_location_type
const (
	ALL_AudiencePackageGetV2DataAudiencePackagesAudienceLocationType     AudiencePackageGetV2DataAudiencePackagesAudienceLocationType = "ALL"
	CURRENT_AudiencePackageGetV2DataAudiencePackagesAudienceLocationType AudiencePackageGetV2DataAudiencePackagesAudienceLocationType = "CURRENT"
	HOME_AudiencePackageGetV2DataAudiencePackagesAudienceLocationType    AudiencePackageGetV2DataAudiencePackagesAudienceLocationType = "HOME"
	TRAVEL_AudiencePackageGetV2DataAudiencePackagesAudienceLocationType  AudiencePackageGetV2DataAudiencePackagesAudienceLocationType = "TRAVEL"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceLocationType enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceLocationTypeEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceLocationTypeFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceLocationTypeFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceLocationType, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceLocationType: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_location_type value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceLocationType) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceLocationType {
	return &v
}
