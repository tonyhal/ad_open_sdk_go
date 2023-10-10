/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2SuperiorPopularityType
type AudiencePackageCreateV2SuperiorPopularityType string

// List of audience_package_create_v2_superior_popularity_type
const (
	GAME_AudiencePackageCreateV2SuperiorPopularityType AudiencePackageCreateV2SuperiorPopularityType = "GAME"
	NONE_AudiencePackageCreateV2SuperiorPopularityType AudiencePackageCreateV2SuperiorPopularityType = "NONE"
)

// All allowed values of AudiencePackageCreateV2SuperiorPopularityType enum
var AllowedAudiencePackageCreateV2SuperiorPopularityTypeEnumValues = []AudiencePackageCreateV2SuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewAudiencePackageCreateV2SuperiorPopularityTypeFromValue returns a pointer to a valid AudiencePackageCreateV2SuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2SuperiorPopularityTypeFromValue(v string) (*AudiencePackageCreateV2SuperiorPopularityType, error) {
	ev := AudiencePackageCreateV2SuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2SuperiorPopularityType: valid values are %v", v, AllowedAudiencePackageCreateV2SuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2SuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2SuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_superior_popularity_type value
func (v AudiencePackageCreateV2SuperiorPopularityType) Ptr() *AudiencePackageCreateV2SuperiorPopularityType {
	return &v
}
