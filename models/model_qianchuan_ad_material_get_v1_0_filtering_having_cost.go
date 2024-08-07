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

// QianchuanAdMaterialGetV10FilteringHavingCost
type QianchuanAdMaterialGetV10FilteringHavingCost string

// List of qianchuan_ad_material_get_v1.0_filtering_having_cost
const (
	ALL_QianchuanAdMaterialGetV10FilteringHavingCost QianchuanAdMaterialGetV10FilteringHavingCost = "ALL"
	YES_QianchuanAdMaterialGetV10FilteringHavingCost QianchuanAdMaterialGetV10FilteringHavingCost = "YES"
)

// All allowed values of QianchuanAdMaterialGetV10FilteringHavingCost enum
var AllowedQianchuanAdMaterialGetV10FilteringHavingCostEnumValues = []QianchuanAdMaterialGetV10FilteringHavingCost{
	"ALL",
	"YES",
}

// NewQianchuanAdMaterialGetV10FilteringHavingCostFromValue returns a pointer to a valid QianchuanAdMaterialGetV10FilteringHavingCost
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdMaterialGetV10FilteringHavingCostFromValue(v string) (*QianchuanAdMaterialGetV10FilteringHavingCost, error) {
	ev := QianchuanAdMaterialGetV10FilteringHavingCost(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdMaterialGetV10FilteringHavingCost: valid values are %v", v, AllowedQianchuanAdMaterialGetV10FilteringHavingCostEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdMaterialGetV10FilteringHavingCost) IsValid() bool {
	for _, existing := range AllowedQianchuanAdMaterialGetV10FilteringHavingCostEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_material_get_v1.0_filtering_having_cost value
func (v QianchuanAdMaterialGetV10FilteringHavingCost) Ptr() *QianchuanAdMaterialGetV10FilteringHavingCost {
	return &v
}