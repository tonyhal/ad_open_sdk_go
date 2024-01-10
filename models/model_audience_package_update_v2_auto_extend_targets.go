/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2AutoExtendTargets
type AudiencePackageUpdateV2AutoExtendTargets string

// List of audience_package_update_v2_auto_extend_targets
const (
	GENDER_AudiencePackageUpdateV2AutoExtendTargets          AudiencePackageUpdateV2AutoExtendTargets = "GENDER"
	REGION_AudiencePackageUpdateV2AutoExtendTargets          AudiencePackageUpdateV2AutoExtendTargets = "REGION"
	AD_TAG_AudiencePackageUpdateV2AutoExtendTargets          AudiencePackageUpdateV2AutoExtendTargets = "AD_TAG"
	INTEREST_TAG_AudiencePackageUpdateV2AutoExtendTargets    AudiencePackageUpdateV2AutoExtendTargets = "INTEREST_TAG"
	AGE_AudiencePackageUpdateV2AutoExtendTargets             AudiencePackageUpdateV2AutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_AudiencePackageUpdateV2AutoExtendTargets AudiencePackageUpdateV2AutoExtendTargets = "CUSTOM_AUDIENCE"
)

// All allowed values of AudiencePackageUpdateV2AutoExtendTargets enum
var AllowedAudiencePackageUpdateV2AutoExtendTargetsEnumValues = []AudiencePackageUpdateV2AutoExtendTargets{
	"GENDER",
	"REGION",
	"AD_TAG",
	"INTEREST_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
}

// NewAudiencePackageUpdateV2AutoExtendTargetsFromValue returns a pointer to a valid AudiencePackageUpdateV2AutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2AutoExtendTargetsFromValue(v string) (*AudiencePackageUpdateV2AutoExtendTargets, error) {
	ev := AudiencePackageUpdateV2AutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2AutoExtendTargets: valid values are %v", v, AllowedAudiencePackageUpdateV2AutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2AutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2AutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_auto_extend_targets value
func (v AudiencePackageUpdateV2AutoExtendTargets) Ptr() *AudiencePackageUpdateV2AutoExtendTargets {
	return &v
}
