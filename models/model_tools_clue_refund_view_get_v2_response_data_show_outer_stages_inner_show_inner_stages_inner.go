/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerShowInnerStagesInner struct for ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerShowInnerStagesInner
type ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerShowInnerStagesInner struct {
	// 是否符合: 0 不符合, 1 符合
	InnerRefundStatus *int64 `json:"inner_refund_status,omitempty"`
	// \"符合\" ｜ \"未符合\"
	InnerRefundStatusText *string                                                                                      `json:"inner_refund_status_text,omitempty"`
	ShowTimeDetails       *ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerShowInnerStagesInnerShowTimeDetails `json:"show_time_details,omitempty"`
	// true | false,是否显示小问号，飞鱼SDK前端所需字段，传入即可
	ShowTip *bool `json:"show_tip,omitempty"`
	// \"为保证首次跟进时间充裕，首通的12小时为可外呼时间段（7:00!23:00）的完整12小时\" ，飞鱼SDK前端所需字段，传入即可
	ShowTipContent *string `json:"show_tip_content,omitempty"`
	// 在 0-12小时内 拨打电话为未接通 | 在 0-24小时内 拨打 3 次 电话为未接通
	Title *string `json:"title,omitempty"`
}
