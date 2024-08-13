/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseItemListV10FilterItemType
type EnterpriseItemListV10FilterItemType string

// List of enterprise_item_list_v1.0_filter_item_type
const (
	ITEM_AD_EnterpriseItemListV10FilterItemType      EnterpriseItemListV10FilterItemType = "ITEM_AD"
	ITEM_CONTENT_EnterpriseItemListV10FilterItemType EnterpriseItemListV10FilterItemType = "ITEM_CONTENT"
)

// All allowed values of EnterpriseItemListV10FilterItemType enum
var AllowedEnterpriseItemListV10FilterItemTypeEnumValues = []EnterpriseItemListV10FilterItemType{
	"ITEM_AD",
	"ITEM_CONTENT",
}

// NewEnterpriseItemListV10FilterItemTypeFromValue returns a pointer to a valid EnterpriseItemListV10FilterItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseItemListV10FilterItemTypeFromValue(v string) (*EnterpriseItemListV10FilterItemType, error) {
	ev := EnterpriseItemListV10FilterItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseItemListV10FilterItemType: valid values are %v", v, AllowedEnterpriseItemListV10FilterItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseItemListV10FilterItemType) IsValid() bool {
	for _, existing := range AllowedEnterpriseItemListV10FilterItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_item_list_v1.0_filter_item_type value
func (v EnterpriseItemListV10FilterItemType) Ptr() *EnterpriseItemListV10FilterItemType {
	return &v
}
