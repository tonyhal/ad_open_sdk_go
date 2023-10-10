/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionListV10OrderType
type QianchuanUniPromotionListV10OrderType string

// List of qianchuan_uni_promotion_list_v1.0_order_type
const (
	ASC_QianchuanUniPromotionListV10OrderType  QianchuanUniPromotionListV10OrderType = "ASC"
	DESC_QianchuanUniPromotionListV10OrderType QianchuanUniPromotionListV10OrderType = "DESC"
)

// All allowed values of QianchuanUniPromotionListV10OrderType enum
var AllowedQianchuanUniPromotionListV10OrderTypeEnumValues = []QianchuanUniPromotionListV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanUniPromotionListV10OrderTypeFromValue returns a pointer to a valid QianchuanUniPromotionListV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10OrderTypeFromValue(v string) (*QianchuanUniPromotionListV10OrderType, error) {
	ev := QianchuanUniPromotionListV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10OrderType: valid values are %v", v, AllowedQianchuanUniPromotionListV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_order_type value
func (v QianchuanUniPromotionListV10OrderType) Ptr() *QianchuanUniPromotionListV10OrderType {
	return &v
}
