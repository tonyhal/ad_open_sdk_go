/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionListV10Fields
type QianchuanUniPromotionListV10Fields string

// List of qianchuan_uni_promotion_list_v1.0_fields
const (
	STAT_COST_QianchuanUniPromotionListV10Fields                         QianchuanUniPromotionListV10Fields = "stat_cost"
	TOTAL_COST_PER_PAY_ORDER_FOR_ROI2_QianchuanUniPromotionListV10Fields QianchuanUniPromotionListV10Fields = "total_cost_per_pay_order_for_roi2"
	TOTAL_PAY_ORDER_COUNT_FOR_ROI2_QianchuanUniPromotionListV10Fields    QianchuanUniPromotionListV10Fields = "total_pay_order_count_for_roi2"
	TOTAL_PAY_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10Fields      QianchuanUniPromotionListV10Fields = "total_pay_order_gmv_for_roi2"
	TOTAL_PREPAY_AND_PAY_ORDER_ROI2_QianchuanUniPromotionListV10Fields   QianchuanUniPromotionListV10Fields = "total_prepay_and_pay_order_roi2"
	TOTAL_PREPAY_ORDER_COUNT_FOR_ROI2_QianchuanUniPromotionListV10Fields QianchuanUniPromotionListV10Fields = "total_prepay_order_count_for_roi2"
	TOTAL_PREPAY_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10Fields   QianchuanUniPromotionListV10Fields = "total_prepay_order_gmv_for_roi2"
)

// All allowed values of QianchuanUniPromotionListV10Fields enum
var AllowedQianchuanUniPromotionListV10FieldsEnumValues = []QianchuanUniPromotionListV10Fields{
	"stat_cost",
	"total_cost_per_pay_order_for_roi2",
	"total_pay_order_count_for_roi2",
	"total_pay_order_gmv_for_roi2",
	"total_prepay_and_pay_order_roi2",
	"total_prepay_order_count_for_roi2",
	"total_prepay_order_gmv_for_roi2",
}

// NewQianchuanUniPromotionListV10FieldsFromValue returns a pointer to a valid QianchuanUniPromotionListV10Fields
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10FieldsFromValue(v string) (*QianchuanUniPromotionListV10Fields, error) {
	ev := QianchuanUniPromotionListV10Fields(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10Fields: valid values are %v", v, AllowedQianchuanUniPromotionListV10FieldsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10Fields) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10FieldsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_fields value
func (v QianchuanUniPromotionListV10Fields) Ptr() *QianchuanUniPromotionListV10Fields {
	return &v
}
