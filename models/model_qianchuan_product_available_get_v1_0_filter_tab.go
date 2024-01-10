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

// QianchuanProductAvailableGetV10FilterTab
type QianchuanProductAvailableGetV10FilterTab string

// List of qianchuan_product_available_get_v1.0_filter_tab
const (
	SEPARATE_NEW_PRODUCT_ALL_QianchuanProductAvailableGetV10FilterTab           QianchuanProductAvailableGetV10FilterTab = "SEPARATE_NEW_PRODUCT_ALL"
	SEPARATE_NEW_PRODUCT_TO_BREAK_TEN_QianchuanProductAvailableGetV10FilterTab  QianchuanProductAvailableGetV10FilterTab = "SEPARATE_NEW_PRODUCT_TO_BREAK_TEN"
	SEPARATE_NEW_PRODUCT_TO_BREAK_ZERO_QianchuanProductAvailableGetV10FilterTab QianchuanProductAvailableGetV10FilterTab = "SEPARATE_NEW_PRODUCT_TO_BREAK_ZERO"
)

// All allowed values of QianchuanProductAvailableGetV10FilterTab enum
var AllowedQianchuanProductAvailableGetV10FilterTabEnumValues = []QianchuanProductAvailableGetV10FilterTab{
	"SEPARATE_NEW_PRODUCT_ALL",
	"SEPARATE_NEW_PRODUCT_TO_BREAK_TEN",
	"SEPARATE_NEW_PRODUCT_TO_BREAK_ZERO",
}

// NewQianchuanProductAvailableGetV10FilterTabFromValue returns a pointer to a valid QianchuanProductAvailableGetV10FilterTab
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAvailableGetV10FilterTabFromValue(v string) (*QianchuanProductAvailableGetV10FilterTab, error) {
	ev := QianchuanProductAvailableGetV10FilterTab(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAvailableGetV10FilterTab: valid values are %v", v, AllowedQianchuanProductAvailableGetV10FilterTabEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAvailableGetV10FilterTab) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAvailableGetV10FilterTabEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_available_get_v1.0_filter_tab value
func (v QianchuanProductAvailableGetV10FilterTab) Ptr() *QianchuanProductAvailableGetV10FilterTab {
	return &v
}
