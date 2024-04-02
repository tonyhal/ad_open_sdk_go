/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdRegionUpdateV10LocationType
type QianchuanAdRegionUpdateV10LocationType string

// List of qianchuan_ad_region_update_v1.0_location_type
const (
	ALL_QianchuanAdRegionUpdateV10LocationType     QianchuanAdRegionUpdateV10LocationType = "ALL"
	CURRENT_QianchuanAdRegionUpdateV10LocationType QianchuanAdRegionUpdateV10LocationType = "CURRENT"
	HOME_QianchuanAdRegionUpdateV10LocationType    QianchuanAdRegionUpdateV10LocationType = "HOME"
	TRAVEL_QianchuanAdRegionUpdateV10LocationType  QianchuanAdRegionUpdateV10LocationType = "TRAVEL"
)

// All allowed values of QianchuanAdRegionUpdateV10LocationType enum
var AllowedQianchuanAdRegionUpdateV10LocationTypeEnumValues = []QianchuanAdRegionUpdateV10LocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewQianchuanAdRegionUpdateV10LocationTypeFromValue returns a pointer to a valid QianchuanAdRegionUpdateV10LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRegionUpdateV10LocationTypeFromValue(v string) (*QianchuanAdRegionUpdateV10LocationType, error) {
	ev := QianchuanAdRegionUpdateV10LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRegionUpdateV10LocationType: valid values are %v", v, AllowedQianchuanAdRegionUpdateV10LocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRegionUpdateV10LocationType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRegionUpdateV10LocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_region_update_v1.0_location_type value
func (v QianchuanAdRegionUpdateV10LocationType) Ptr() *QianchuanAdRegionUpdateV10LocationType {
	return &v
}
