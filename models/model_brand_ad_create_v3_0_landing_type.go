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

// BrandAdCreateV30LandingType
type BrandAdCreateV30LandingType string

// List of brand_ad_create_v3.0_landing_type
const (
	BRAND_CONTENT_BrandAdCreateV30LandingType BrandAdCreateV30LandingType = "BRAND_CONTENT"
)

// All allowed values of BrandAdCreateV30LandingType enum
var AllowedBrandAdCreateV30LandingTypeEnumValues = []BrandAdCreateV30LandingType{
	"BRAND_CONTENT",
}

// NewBrandAdCreateV30LandingTypeFromValue returns a pointer to a valid BrandAdCreateV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30LandingTypeFromValue(v string) (*BrandAdCreateV30LandingType, error) {
	ev := BrandAdCreateV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30LandingType: valid values are %v", v, AllowedBrandAdCreateV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30LandingType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_landing_type value
func (v BrandAdCreateV30LandingType) Ptr() *BrandAdCreateV30LandingType {
	return &v
}
