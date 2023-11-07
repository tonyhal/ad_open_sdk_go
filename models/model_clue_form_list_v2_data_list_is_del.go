/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormListV2DataListIsDel
type ClueFormListV2DataListIsDel string

// List of clue_form_list_v2_data_list_is_del
const (
	Enum_0_ClueFormListV2DataListIsDel ClueFormListV2DataListIsDel = "0"
	Enum_1_ClueFormListV2DataListIsDel ClueFormListV2DataListIsDel = "1"
)

// All allowed values of ClueFormListV2DataListIsDel enum
var AllowedClueFormListV2DataListIsDelEnumValues = []ClueFormListV2DataListIsDel{
	"0",
	"1",
}

// NewClueFormListV2DataListIsDelFromValue returns a pointer to a valid ClueFormListV2DataListIsDel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormListV2DataListIsDelFromValue(v string) (*ClueFormListV2DataListIsDel, error) {
	ev := ClueFormListV2DataListIsDel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormListV2DataListIsDel: valid values are %v", v, AllowedClueFormListV2DataListIsDelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormListV2DataListIsDel) IsValid() bool {
	for _, existing := range AllowedClueFormListV2DataListIsDelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_list_v2_data_list_is_del value
func (v ClueFormListV2DataListIsDel) Ptr() *ClueFormListV2DataListIsDel {
	return &v
}
