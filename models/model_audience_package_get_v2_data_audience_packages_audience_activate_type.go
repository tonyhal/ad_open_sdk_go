/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceActivateType
type AudiencePackageGetV2DataAudiencePackagesAudienceActivateType string

// List of audience_package_get_v2_data_audience_packages_audience_activate_type
const (
	ONE_MONTH_2_THREE_MONTH_AudiencePackageGetV2DataAudiencePackagesAudienceActivateType AudiencePackageGetV2DataAudiencePackagesAudienceActivateType = "ONE_MONTH_2_THREE_MONTH"
	THREE_MONTH_EAILIER_AudiencePackageGetV2DataAudiencePackagesAudienceActivateType     AudiencePackageGetV2DataAudiencePackagesAudienceActivateType = "THREE_MONTH_EAILIER"
	UNLIMITED_AudiencePackageGetV2DataAudiencePackagesAudienceActivateType               AudiencePackageGetV2DataAudiencePackagesAudienceActivateType = "UNLIMITED"
	WITH_IN_A_MONTH_AudiencePackageGetV2DataAudiencePackagesAudienceActivateType         AudiencePackageGetV2DataAudiencePackagesAudienceActivateType = "WITH_IN_A_MONTH"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceActivateType enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActivateTypeEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceActivateType{
	"ONE_MONTH_2_THREE_MONTH",
	"THREE_MONTH_EAILIER",
	"UNLIMITED",
	"WITH_IN_A_MONTH",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceActivateTypeFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceActivateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceActivateTypeFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceActivateType, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceActivateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceActivateType: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActivateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceActivateType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActivateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_activate_type value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceActivateType) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceActivateType {
	return &v
}
