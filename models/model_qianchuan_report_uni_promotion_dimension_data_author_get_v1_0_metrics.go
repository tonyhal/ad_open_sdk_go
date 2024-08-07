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

// QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics
type QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics string

// List of qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_metrics
const (
	STAT_COST_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics                                    QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "stat_cost"
	TOTAL_COST_PER_PAY_ORDER_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics            QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_cost_per_pay_order_for_roi2"
	TOTAL_PAY_ORDER_COUNT_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics               QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_pay_order_count_for_roi2"
	TOTAL_PAY_ORDER_COUPON_AMOUNT_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics       QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_pay_order_coupon_amount_for_roi2"
	TOTAL_PAY_ORDER_GMV_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics                 QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_pay_order_gmv_for_roi2"
	TOTAL_PREPAY_AND_PAY_ORDER_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics              QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_prepay_and_pay_order_roi2"
	TOTAL_PREPAY_ORDER_COUNT_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics            QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_prepay_order_count_for_roi2"
	TOTAL_PREPAY_ORDER_GMV_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics              QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_prepay_order_gmv_for_roi2"
	TOTAL_UNFINISHED_ESTIMATE_ORDER_GMV_FOR_ROI2_QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics = "total_unfinished_estimate_order_gmv_for_roi2"
)

// All allowed values of QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics enum
var AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10MetricsEnumValues = []QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics{
	"stat_cost",
	"total_cost_per_pay_order_for_roi2",
	"total_pay_order_count_for_roi2",
	"total_pay_order_coupon_amount_for_roi2",
	"total_pay_order_gmv_for_roi2",
	"total_prepay_and_pay_order_roi2",
	"total_prepay_order_count_for_roi2",
	"total_prepay_order_gmv_for_roi2",
	"total_unfinished_estimate_order_gmv_for_roi2",
}

// NewQianchuanReportUniPromotionDimensionDataAuthorGetV10MetricsFromValue returns a pointer to a valid QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionDimensionDataAuthorGetV10MetricsFromValue(v string) (*QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics, error) {
	ev := QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics: valid values are %v", v, AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10MetricsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10MetricsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_metrics value
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics) Ptr() *QianchuanReportUniPromotionDimensionDataAuthorGetV10Metrics {
	return &v
}
