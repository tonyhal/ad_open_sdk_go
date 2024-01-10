/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaExpGetV2ResponseDataDataArrayInner struct for ReportRtaExpGetV2ResponseDataDataArrayInner
type ReportRtaExpGetV2ResponseDataDataArrayInner struct {
	// 展现数据-点击数 当头条用户点击广告素材时，触发点击事件，该事件被认为是一次有效的广告点击
	Click *int64 `json:"click,omitempty"`
	// 转化数据-转化数 将转化数记录在转化事件发生的时间上。建议广告主考核成本时参考“转化数据（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点
	Convert *int64 `json:"convert,omitempty"`
	// 展现数据-总花费 表示广告在投放期内的预估花费金额,当天数据可能会有波动，次日稳定
	Cost *float64 `json:"cost,omitempty"`
	// 数据统计日期，格式YYYYMMDD
	Date *string `json:"date,omitempty"`
	// 展现数据-展示数 广告展示给用户的次数。计算方式：经平台判定有效且被计费的展示次数
	Show *int64 `json:"show,omitempty"`
	// 联合实验策略，请求入参
	Strategy *string `json:"strategy,omitempty"`
	// 竞胜率。 竞胜率=竞胜数/参竞数，代表广告主参竞请求的胜出比例
	WinRatio *float64 `json:"win_ratio,omitempty"`
}
