/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandDataV30ResponseData
type ReportBrandDataV30ResponseData struct {
	//
	DataReports []*ReportBrandDataV30ResponseDataDataReportsInner `json:"data_reports,omitempty"`
	PageInfo    *ReportBrandDataV30ResponseDataPageInfo           `json:"page_info,omitempty"`
}
