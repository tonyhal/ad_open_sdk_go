/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderReportV30ResponseData
type DouplusOrderReportV30ResponseData struct {
	//
	OrderMetrics []*DouplusOrderReportV30ResponseDataOrderMetricsInner `json:"order_metrics,omitempty"`
	PageInfo     *DouplusOrderListV30ResponseDataPageInfo              `json:"page_info,omitempty"`
}
