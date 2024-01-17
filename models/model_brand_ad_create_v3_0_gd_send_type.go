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

// BrandAdCreateV30GdSendType
type BrandAdCreateV30GdSendType string

// List of brand_ad_create_v3.0_gd_send_type
const (
	CTR_BrandAdCreateV30GdSendType         BrandAdCreateV30GdSendType = "CTR"
	FANS_INCR_BrandAdCreateV30GdSendType   BrandAdCreateV30GdSendType = "FANS_INCR"
	FORM_BrandAdCreateV30GdSendType        BrandAdCreateV30GdSendType = "FORM"
	HOISTING_BrandAdCreateV30GdSendType    BrandAdCreateV30GdSendType = "HOISTING"
	INTERACTIVE_BrandAdCreateV30GdSendType BrandAdCreateV30GdSendType = "INTERACTIVE"
	PLANT_GRASS_BrandAdCreateV30GdSendType BrandAdCreateV30GdSendType = "PLANT_GRASS"
	SEQUENCE_BrandAdCreateV30GdSendType    BrandAdCreateV30GdSendType = "SEQUENCE"
)

// All allowed values of BrandAdCreateV30GdSendType enum
var AllowedBrandAdCreateV30GdSendTypeEnumValues = []BrandAdCreateV30GdSendType{
	"CTR",
	"FANS_INCR",
	"FORM",
	"HOISTING",
	"INTERACTIVE",
	"PLANT_GRASS",
	"SEQUENCE",
}

// NewBrandAdCreateV30GdSendTypeFromValue returns a pointer to a valid BrandAdCreateV30GdSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30GdSendTypeFromValue(v string) (*BrandAdCreateV30GdSendType, error) {
	ev := BrandAdCreateV30GdSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30GdSendType: valid values are %v", v, AllowedBrandAdCreateV30GdSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30GdSendType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30GdSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_gd_send_type value
func (v BrandAdCreateV30GdSendType) Ptr() *BrandAdCreateV30GdSendType {
	return &v
}