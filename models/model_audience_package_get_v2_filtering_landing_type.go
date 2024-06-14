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

// AudiencePackageGetV2FilteringLandingType
type AudiencePackageGetV2FilteringLandingType string

// List of audience_package_get_v2_filtering_landing_type
const (
	APP_AudiencePackageGetV2FilteringLandingType         AudiencePackageGetV2FilteringLandingType = "APP"
	APP_ANDROID_AudiencePackageGetV2FilteringLandingType AudiencePackageGetV2FilteringLandingType = "APP_ANDROID"
	APP_IOS_AudiencePackageGetV2FilteringLandingType     AudiencePackageGetV2FilteringLandingType = "APP_IOS"
	ARTICLE_AudiencePackageGetV2FilteringLandingType     AudiencePackageGetV2FilteringLandingType = "ARTICLE"
	AWEME_AudiencePackageGetV2FilteringLandingType       AudiencePackageGetV2FilteringLandingType = "AWEME"
	DPA_AudiencePackageGetV2FilteringLandingType         AudiencePackageGetV2FilteringLandingType = "DPA"
	EXTERNAL_AudiencePackageGetV2FilteringLandingType    AudiencePackageGetV2FilteringLandingType = "EXTERNAL"
	GOODS_AudiencePackageGetV2FilteringLandingType       AudiencePackageGetV2FilteringLandingType = "GOODS"
	LIVE_AudiencePackageGetV2FilteringLandingType        AudiencePackageGetV2FilteringLandingType = "LIVE"
	MICRO_GAME_AudiencePackageGetV2FilteringLandingType  AudiencePackageGetV2FilteringLandingType = "MICRO_GAME"
	QUICK_APP_AudiencePackageGetV2FilteringLandingType   AudiencePackageGetV2FilteringLandingType = "QUICK_APP"
	SHOP_AudiencePackageGetV2FilteringLandingType        AudiencePackageGetV2FilteringLandingType = "SHOP"
	STORE_AudiencePackageGetV2FilteringLandingType       AudiencePackageGetV2FilteringLandingType = "STORE"
)

// All allowed values of AudiencePackageGetV2FilteringLandingType enum
var AllowedAudiencePackageGetV2FilteringLandingTypeEnumValues = []AudiencePackageGetV2FilteringLandingType{
	"APP",
	"APP_ANDROID",
	"APP_IOS",
	"ARTICLE",
	"AWEME",
	"DPA",
	"EXTERNAL",
	"GOODS",
	"LIVE",
	"MICRO_GAME",
	"QUICK_APP",
	"SHOP",
	"STORE",
}

// NewAudiencePackageGetV2FilteringLandingTypeFromValue returns a pointer to a valid AudiencePackageGetV2FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2FilteringLandingTypeFromValue(v string) (*AudiencePackageGetV2FilteringLandingType, error) {
	ev := AudiencePackageGetV2FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2FilteringLandingType: valid values are %v", v, AllowedAudiencePackageGetV2FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_filtering_landing_type value
func (v AudiencePackageGetV2FilteringLandingType) Ptr() *AudiencePackageGetV2FilteringLandingType {
	return &v
}
