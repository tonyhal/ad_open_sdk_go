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

// BrandAdCreateV30AppOrigin
type BrandAdCreateV30AppOrigin string

// List of brand_ad_create_v3.0_app_origin
const (
	AWEME_BrandAdCreateV30AppOrigin BrandAdCreateV30AppOrigin = "AWEME"
)

// All allowed values of BrandAdCreateV30AppOrigin enum
var AllowedBrandAdCreateV30AppOriginEnumValues = []BrandAdCreateV30AppOrigin{
	"AWEME",
}

// NewBrandAdCreateV30AppOriginFromValue returns a pointer to a valid BrandAdCreateV30AppOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AppOriginFromValue(v string) (*BrandAdCreateV30AppOrigin, error) {
	ev := BrandAdCreateV30AppOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AppOrigin: valid values are %v", v, AllowedBrandAdCreateV30AppOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AppOrigin) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AppOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_app_origin value
func (v BrandAdCreateV30AppOrigin) Ptr() *BrandAdCreateV30AppOrigin {
	return &v
}