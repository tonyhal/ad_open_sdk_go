/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomCreativeGetV30ResponseData
type ReportCustomCreativeGetV30ResponseData struct {
	PageInfo *ReportCustomCreativeGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Rows []*ReportCustomCreativeGetV30ResponseDataRowsInner `json:"rows"`
	// 指标汇总数据
	TotalMetrics map[string]string `json:"total_metrics"`
}
