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

// AudiencePackageGetV2DataAudiencePackagesAudienceGender
type AudiencePackageGetV2DataAudiencePackagesAudienceGender string

// List of audience_package_get_v2_data_audience_packages_audience_gender
const (
	GENDER_FEMALE_AudiencePackageGetV2DataAudiencePackagesAudienceGender    AudiencePackageGetV2DataAudiencePackagesAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_AudiencePackageGetV2DataAudiencePackagesAudienceGender      AudiencePackageGetV2DataAudiencePackagesAudienceGender = "GENDER_MALE"
	GENDER_UNLIMITED_AudiencePackageGetV2DataAudiencePackagesAudienceGender AudiencePackageGetV2DataAudiencePackagesAudienceGender = "GENDER_UNLIMITED"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceGender enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceGenderEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"GENDER_UNLIMITED",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceGenderFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceGenderFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceGender, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceGender: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceGender) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_gender value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceGender) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceGender {
	return &v
}
