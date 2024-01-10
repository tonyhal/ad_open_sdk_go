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

// BrandQueryStockV30AudienceInfoAc
type BrandQueryStockV30AudienceInfoAc string

// List of brand_query_stock_v3.0_audience_info_ac
const (
	UNLIMITED_BrandQueryStockV30AudienceInfoAc BrandQueryStockV30AudienceInfoAc = "UNLIMITED"
	WIFI_BrandQueryStockV30AudienceInfoAc      BrandQueryStockV30AudienceInfoAc = "WIFI"
)

// All allowed values of BrandQueryStockV30AudienceInfoAc enum
var AllowedBrandQueryStockV30AudienceInfoAcEnumValues = []BrandQueryStockV30AudienceInfoAc{
	"UNLIMITED",
	"WIFI",
}

// NewBrandQueryStockV30AudienceInfoAcFromValue returns a pointer to a valid BrandQueryStockV30AudienceInfoAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AudienceInfoAcFromValue(v string) (*BrandQueryStockV30AudienceInfoAc, error) {
	ev := BrandQueryStockV30AudienceInfoAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AudienceInfoAc: valid values are %v", v, AllowedBrandQueryStockV30AudienceInfoAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AudienceInfoAc) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AudienceInfoAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_audience_info_ac value
func (v BrandQueryStockV30AudienceInfoAc) Ptr() *BrandQueryStockV30AudienceInfoAc {
	return &v
}
