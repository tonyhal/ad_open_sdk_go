/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAdMetricsV2ResponseDataStatsValue struct for CreativeAdMetricsV2ResponseDataStatsValue
type CreativeAdMetricsV2ResponseDataStatsValue struct {
	// 广告计划 ID
	AdId *int64 `json:"ad_id,omitempty"`
	// 派生消耗（单位：元）
	DerivedCost *float64 `json:"derived_cost,omitempty"`
	// 派生消耗占比
	DerivedPercent *float64 `json:"derived_percent,omitempty"`
	// 总计消耗（单位：元）
	TotalCost *float64 `json:"total_cost,omitempty"`
}
