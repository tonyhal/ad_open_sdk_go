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

// QianchuanAdCreateV10CreativeAutoGenerate
type QianchuanAdCreateV10CreativeAutoGenerate int64

// List of qianchuan_ad_create_v1.0_creative_auto_generate
const (
	Enum_0_QianchuanAdCreateV10CreativeAutoGenerate QianchuanAdCreateV10CreativeAutoGenerate = 0
	Enum_1_QianchuanAdCreateV10CreativeAutoGenerate QianchuanAdCreateV10CreativeAutoGenerate = 1
)

// All allowed values of QianchuanAdCreateV10CreativeAutoGenerate enum
var AllowedQianchuanAdCreateV10CreativeAutoGenerateEnumValues = []QianchuanAdCreateV10CreativeAutoGenerate{
	0,
	1,
}

// NewQianchuanAdCreateV10CreativeAutoGenerateFromValue returns a pointer to a valid QianchuanAdCreateV10CreativeAutoGenerate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10CreativeAutoGenerateFromValue(v int64) (*QianchuanAdCreateV10CreativeAutoGenerate, error) {
	ev := QianchuanAdCreateV10CreativeAutoGenerate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10CreativeAutoGenerate: valid values are %v", v, AllowedQianchuanAdCreateV10CreativeAutoGenerateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10CreativeAutoGenerate) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10CreativeAutoGenerateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_creative_auto_generate value
func (v QianchuanAdCreateV10CreativeAutoGenerate) Ptr() *QianchuanAdCreateV10CreativeAutoGenerate {
	return &v
}
