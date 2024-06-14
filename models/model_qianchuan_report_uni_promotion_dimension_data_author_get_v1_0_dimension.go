/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension
type QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension string

// List of qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_dimension
const (
	TIME_GRANULARITY_DAILY_QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension  QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_HOURLY_QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension = "TIME_GRANULARITY_HOURLY"
)

// All allowed values of QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension enum
var AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10DimensionEnumValues = []QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_HOURLY",
}

// NewQianchuanReportUniPromotionDimensionDataAuthorGetV10DimensionFromValue returns a pointer to a valid QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionDimensionDataAuthorGetV10DimensionFromValue(v string) (*QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension, error) {
	ev := QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension: valid values are %v", v, AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10DimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10DimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_dimension value
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension) Ptr() *QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension {
	return &v
}
