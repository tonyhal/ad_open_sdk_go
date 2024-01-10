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

// AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType
type AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType string

// List of adlab_group_detail_v3.0_data_data_ad_info_product_info_supplements_supplement_type
const (
	SUPPLEMENT_AUXILIARY_AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType = "SUPPLEMENT_AUXILIARY"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType{
	"SUPPLEMENT_AUXILIARY",
}

// NewAdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_product_info_supplements_supplement_type value
func (v AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType) Ptr() *AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType {
	return &v
}
