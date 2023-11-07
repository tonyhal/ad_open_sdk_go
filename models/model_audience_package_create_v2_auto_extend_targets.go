/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2AutoExtendTargets
type AudiencePackageCreateV2AutoExtendTargets string

// List of audience_package_create_v2_auto_extend_targets
const (
	AD_TAG_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "AD_TAG"
	GENDER_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "GENDER"
	AGE_AudiencePackageCreateV2AutoExtendTargets             AudiencePackageCreateV2AutoExtendTargets = "AGE"
	REGION_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "REGION"
	CUSTOM_AUDIENCE_AudiencePackageCreateV2AutoExtendTargets AudiencePackageCreateV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	INTEREST_TAG_AudiencePackageCreateV2AutoExtendTargets    AudiencePackageCreateV2AutoExtendTargets = "INTEREST_TAG"
)

// All allowed values of AudiencePackageCreateV2AutoExtendTargets enum
var AllowedAudiencePackageCreateV2AutoExtendTargetsEnumValues = []AudiencePackageCreateV2AutoExtendTargets{
	"AD_TAG",
	"GENDER",
	"AGE",
	"REGION",
	"CUSTOM_AUDIENCE",
	"INTEREST_TAG",
}

// NewAudiencePackageCreateV2AutoExtendTargetsFromValue returns a pointer to a valid AudiencePackageCreateV2AutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AutoExtendTargetsFromValue(v string) (*AudiencePackageCreateV2AutoExtendTargets, error) {
	ev := AudiencePackageCreateV2AutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2AutoExtendTargets: valid values are %v", v, AllowedAudiencePackageCreateV2AutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2AutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_auto_extend_targets value
func (v AudiencePackageCreateV2AutoExtendTargets) Ptr() *AudiencePackageCreateV2AutoExtendTargets {
	return &v
}
