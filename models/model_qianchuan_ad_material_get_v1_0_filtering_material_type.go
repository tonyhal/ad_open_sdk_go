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

// QianchuanAdMaterialGetV10FilteringMaterialType
type QianchuanAdMaterialGetV10FilteringMaterialType string

// List of qianchuan_ad_material_get_v1.0_filtering_material_type
const (
	IMAGE_QianchuanAdMaterialGetV10FilteringMaterialType     QianchuanAdMaterialGetV10FilteringMaterialType = "IMAGE"
	TITLE_QianchuanAdMaterialGetV10FilteringMaterialType     QianchuanAdMaterialGetV10FilteringMaterialType = "TITLE"
	LIVE_ROOM_QianchuanAdMaterialGetV10FilteringMaterialType QianchuanAdMaterialGetV10FilteringMaterialType = "LIVE_ROOM"
	VIDEO_QianchuanAdMaterialGetV10FilteringMaterialType     QianchuanAdMaterialGetV10FilteringMaterialType = "VIDEO"
)

// All allowed values of QianchuanAdMaterialGetV10FilteringMaterialType enum
var AllowedQianchuanAdMaterialGetV10FilteringMaterialTypeEnumValues = []QianchuanAdMaterialGetV10FilteringMaterialType{
	"IMAGE",
	"TITLE",
	"LIVE_ROOM",
	"VIDEO",
}

// NewQianchuanAdMaterialGetV10FilteringMaterialTypeFromValue returns a pointer to a valid QianchuanAdMaterialGetV10FilteringMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdMaterialGetV10FilteringMaterialTypeFromValue(v string) (*QianchuanAdMaterialGetV10FilteringMaterialType, error) {
	ev := QianchuanAdMaterialGetV10FilteringMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdMaterialGetV10FilteringMaterialType: valid values are %v", v, AllowedQianchuanAdMaterialGetV10FilteringMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdMaterialGetV10FilteringMaterialType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdMaterialGetV10FilteringMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_material_get_v1.0_filtering_material_type value
func (v QianchuanAdMaterialGetV10FilteringMaterialType) Ptr() *QianchuanAdMaterialGetV10FilteringMaterialType {
	return &v
}