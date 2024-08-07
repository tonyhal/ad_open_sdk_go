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

// BrandQueryStockV30AudienceInfoGender
type BrandQueryStockV30AudienceInfoGender string

// List of brand_query_stock_v3.0_audience_info_gender
const (
	FEMALE_BrandQueryStockV30AudienceInfoGender    BrandQueryStockV30AudienceInfoGender = "FEMALE"
	MALE_BrandQueryStockV30AudienceInfoGender      BrandQueryStockV30AudienceInfoGender = "MALE"
	UNLIMITED_BrandQueryStockV30AudienceInfoGender BrandQueryStockV30AudienceInfoGender = "UNLIMITED"
)

// All allowed values of BrandQueryStockV30AudienceInfoGender enum
var AllowedBrandQueryStockV30AudienceInfoGenderEnumValues = []BrandQueryStockV30AudienceInfoGender{
	"FEMALE",
	"MALE",
	"UNLIMITED",
}

// NewBrandQueryStockV30AudienceInfoGenderFromValue returns a pointer to a valid BrandQueryStockV30AudienceInfoGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AudienceInfoGenderFromValue(v string) (*BrandQueryStockV30AudienceInfoGender, error) {
	ev := BrandQueryStockV30AudienceInfoGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AudienceInfoGender: valid values are %v", v, AllowedBrandQueryStockV30AudienceInfoGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AudienceInfoGender) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AudienceInfoGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_audience_info_gender value
func (v BrandQueryStockV30AudienceInfoGender) Ptr() *BrandQueryStockV30AudienceInfoGender {
	return &v
}
