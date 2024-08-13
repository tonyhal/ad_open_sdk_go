/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaCusExpGetV2ResponseDataListInner struct for ReportRtaCusExpGetV2ResponseDataListInner
type ReportRtaCusExpGetV2ResponseDataListInner struct {
	// rta参竞数
	BidCount *int64 `json:"bid_count,omitempty"`
	// 点击数
	ClickCount *int64 `json:"click_count,omitempty"`
	// 转化数
	ConvertCount *int64 `json:"convert_count,omitempty"`
	// 消耗（元）
	Cost *float64 `json:"cost,omitempty"`
	// 时间
	Date *string `json:"date,omitempty"`
	// rta请求数
	RequestCount *int64 `json:"request_count,omitempty"`
	// rta反向联合实验id
	RtaVid *string `json:"rta_vid,omitempty"`
	// 展示数
	ShowCount *int64 `json:"show_count,omitempty"`
	// 竞胜数
	WinCount *int64 `json:"win_count,omitempty"`
}
