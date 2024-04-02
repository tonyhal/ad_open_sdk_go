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

// ReportCustomAsyncTaskCreateV30OrderByType
type ReportCustomAsyncTaskCreateV30OrderByType string

// List of report_custom_async_task_create_v3.0_order_by_type
const (
	ASC_ReportCustomAsyncTaskCreateV30OrderByType  ReportCustomAsyncTaskCreateV30OrderByType = "ASC"
	DESC_ReportCustomAsyncTaskCreateV30OrderByType ReportCustomAsyncTaskCreateV30OrderByType = "DESC"
)

// All allowed values of ReportCustomAsyncTaskCreateV30OrderByType enum
var AllowedReportCustomAsyncTaskCreateV30OrderByTypeEnumValues = []ReportCustomAsyncTaskCreateV30OrderByType{
	"ASC",
	"DESC",
}

// NewReportCustomAsyncTaskCreateV30OrderByTypeFromValue returns a pointer to a valid ReportCustomAsyncTaskCreateV30OrderByType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomAsyncTaskCreateV30OrderByTypeFromValue(v string) (*ReportCustomAsyncTaskCreateV30OrderByType, error) {
	ev := ReportCustomAsyncTaskCreateV30OrderByType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomAsyncTaskCreateV30OrderByType: valid values are %v", v, AllowedReportCustomAsyncTaskCreateV30OrderByTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomAsyncTaskCreateV30OrderByType) IsValid() bool {
	for _, existing := range AllowedReportCustomAsyncTaskCreateV30OrderByTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_async_task_create_v3.0_order_by_type value
func (v ReportCustomAsyncTaskCreateV30OrderByType) Ptr() *ReportCustomAsyncTaskCreateV30OrderByType {
	return &v
}
