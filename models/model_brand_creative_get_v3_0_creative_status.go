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

// BrandCreativeGetV30CreativeStatus
type BrandCreativeGetV30CreativeStatus int64

// List of brand_creative_get_v3.0_creative_status
const (
	Enum_2_BrandCreativeGetV30CreativeStatus  BrandCreativeGetV30CreativeStatus = 2
	Enum_3_BrandCreativeGetV30CreativeStatus  BrandCreativeGetV30CreativeStatus = 3
	Enum_4_BrandCreativeGetV30CreativeStatus  BrandCreativeGetV30CreativeStatus = 4
	Enum_41_BrandCreativeGetV30CreativeStatus BrandCreativeGetV30CreativeStatus = 41
	Enum_0_BrandCreativeGetV30CreativeStatus  BrandCreativeGetV30CreativeStatus = 0
)

// All allowed values of BrandCreativeGetV30CreativeStatus enum
var AllowedBrandCreativeGetV30CreativeStatusEnumValues = []BrandCreativeGetV30CreativeStatus{
	2,
	3,
	4,
	41,
	0,
}

// NewBrandCreativeGetV30CreativeStatusFromValue returns a pointer to a valid BrandCreativeGetV30CreativeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30CreativeStatusFromValue(v int64) (*BrandCreativeGetV30CreativeStatus, error) {
	ev := BrandCreativeGetV30CreativeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30CreativeStatus: valid values are %v", v, AllowedBrandCreativeGetV30CreativeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30CreativeStatus) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30CreativeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_creative_status value
func (v BrandCreativeGetV30CreativeStatus) Ptr() *BrandCreativeGetV30CreativeStatus {
	return &v
}
