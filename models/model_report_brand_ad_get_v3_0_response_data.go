/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandAdGetV30ResponseData
type ReportBrandAdGetV30ResponseData struct {
	// 报表列表
	DataReports []*ReportBrandAdGetV30ResponseDataDataReportsInner `json:"data_reports,omitempty"`
	// 结果总数
	TotalCount *int64 `json:"total_count,omitempty"`
}
