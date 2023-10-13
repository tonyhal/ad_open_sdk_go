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

// AudiencePackageUpdateV2Platform
type AudiencePackageUpdateV2Platform string

// List of audience_package_update_v2_platform
const (
	WAP_AudiencePackageUpdateV2Platform     AudiencePackageUpdateV2Platform = "WAP"
	ANDROID_AudiencePackageUpdateV2Platform AudiencePackageUpdateV2Platform = "ANDROID"
	PC_AudiencePackageUpdateV2Platform      AudiencePackageUpdateV2Platform = "PC"
	IOS_AudiencePackageUpdateV2Platform     AudiencePackageUpdateV2Platform = "IOS"
	IPAD_AudiencePackageUpdateV2Platform    AudiencePackageUpdateV2Platform = "IPAD"
)

// All allowed values of AudiencePackageUpdateV2Platform enum
var AllowedAudiencePackageUpdateV2PlatformEnumValues = []AudiencePackageUpdateV2Platform{
	"WAP",
	"ANDROID",
	"PC",
	"IOS",
	"IPAD",
}

// NewAudiencePackageUpdateV2PlatformFromValue returns a pointer to a valid AudiencePackageUpdateV2Platform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2PlatformFromValue(v string) (*AudiencePackageUpdateV2Platform, error) {
	ev := AudiencePackageUpdateV2Platform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2Platform: valid values are %v", v, AllowedAudiencePackageUpdateV2PlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2Platform) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2PlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_platform value
func (v AudiencePackageUpdateV2Platform) Ptr() *AudiencePackageUpdateV2Platform {
	return &v
}
