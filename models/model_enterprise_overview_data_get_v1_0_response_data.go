/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseOverviewDataGetV10ResponseData
type EnterpriseOverviewDataGetV10ResponseData struct {
	//
	List []*EnterpriseOverviewDataGetV10ResponseDataListInner `json:"list,omitempty"`
	//
	Offset *int64 `json:"offset,omitempty"`
	//
	TotalCount   *int64                                                `json:"total_count,omitempty"`
	TotalMetrics *EnterpriseOverviewDataGetV10ResponseDataTotalMetrics `json:"total_metrics,omitempty"`
	TotalRatio   *EnterpriseOverviewDataGetV10ResponseDataTotalRatio   `json:"total_ratio,omitempty"`
}
