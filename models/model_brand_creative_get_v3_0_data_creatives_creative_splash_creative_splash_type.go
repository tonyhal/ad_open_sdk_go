/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType
type BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType int64

// List of brand_creative_get_v3.0_data_creatives_creative_splash_creative_splash_type
const (
	Enum_2_BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType = 2
	Enum_3_BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType = 3
	Enum_4_BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType = 4
	Enum_5_BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType = 5
	Enum_0_BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType = 0
	Enum_1_BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType = 1
)

// All allowed values of BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType enum
var AllowedBrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashTypeEnumValues = []BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType{
	2,
	3,
	4,
	5,
	0,
	1,
}

// NewBrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashTypeFromValue returns a pointer to a valid BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashTypeFromValue(v int64) (*BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType, error) {
	ev := BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType: valid values are %v", v, AllowedBrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_data_creatives_creative_splash_creative_splash_type value
func (v BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType) Ptr() *BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType {
	return &v
}