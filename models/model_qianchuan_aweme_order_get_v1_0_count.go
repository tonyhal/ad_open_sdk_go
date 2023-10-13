/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderGetV10Count
type QianchuanAwemeOrderGetV10Count int64

// List of qianchuan_aweme_order_get_v1.0_count
const (
	Enum_10_QianchuanAwemeOrderGetV10Count QianchuanAwemeOrderGetV10Count = 10
	Enum_20_QianchuanAwemeOrderGetV10Count QianchuanAwemeOrderGetV10Count = 20
	Enum_50_QianchuanAwemeOrderGetV10Count QianchuanAwemeOrderGetV10Count = 50
)

// All allowed values of QianchuanAwemeOrderGetV10Count enum
var AllowedQianchuanAwemeOrderGetV10CountEnumValues = []QianchuanAwemeOrderGetV10Count{
	10,
	20,
	50,
}

// NewQianchuanAwemeOrderGetV10CountFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10Count
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10CountFromValue(v int64) (*QianchuanAwemeOrderGetV10Count, error) {
	ev := QianchuanAwemeOrderGetV10Count(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10Count: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10CountEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10Count) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10CountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_count value
func (v QianchuanAwemeOrderGetV10Count) Ptr() *QianchuanAwemeOrderGetV10Count {
	return &v
}
