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

// AudiencePackageCreateV2LandingType
type AudiencePackageCreateV2LandingType string

// List of audience_package_create_v2_landing_type
const (
	APP_IOS_AudiencePackageCreateV2LandingType     AudiencePackageCreateV2LandingType = "APP_IOS"
	QUICK_APP_AudiencePackageCreateV2LandingType   AudiencePackageCreateV2LandingType = "QUICK_APP"
	DPA_AudiencePackageCreateV2LandingType         AudiencePackageCreateV2LandingType = "DPA"
	ARTICLE_AudiencePackageCreateV2LandingType     AudiencePackageCreateV2LandingType = "ARTICLE"
	AWEME_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "AWEME"
	LIVE_AudiencePackageCreateV2LandingType        AudiencePackageCreateV2LandingType = "LIVE"
	SHOP_AudiencePackageCreateV2LandingType        AudiencePackageCreateV2LandingType = "SHOP"
	APP_ANDROID_AudiencePackageCreateV2LandingType AudiencePackageCreateV2LandingType = "APP_ANDROID"
	STORE_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "STORE"
	EXTERNAL_AudiencePackageCreateV2LandingType    AudiencePackageCreateV2LandingType = "EXTERNAL"
	MICRO_GAME_AudiencePackageCreateV2LandingType  AudiencePackageCreateV2LandingType = "MICRO_GAME"
	GOODS_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "GOODS"
)

// All allowed values of AudiencePackageCreateV2LandingType enum
var AllowedAudiencePackageCreateV2LandingTypeEnumValues = []AudiencePackageCreateV2LandingType{
	"APP_IOS",
	"QUICK_APP",
	"DPA",
	"ARTICLE",
	"AWEME",
	"LIVE",
	"SHOP",
	"APP_ANDROID",
	"STORE",
	"EXTERNAL",
	"MICRO_GAME",
	"GOODS",
}

// NewAudiencePackageCreateV2LandingTypeFromValue returns a pointer to a valid AudiencePackageCreateV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2LandingTypeFromValue(v string) (*AudiencePackageCreateV2LandingType, error) {
	ev := AudiencePackageCreateV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2LandingType: valid values are %v", v, AllowedAudiencePackageCreateV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2LandingType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_landing_type value
func (v AudiencePackageCreateV2LandingType) Ptr() *AudiencePackageCreateV2LandingType {
	return &v
}
