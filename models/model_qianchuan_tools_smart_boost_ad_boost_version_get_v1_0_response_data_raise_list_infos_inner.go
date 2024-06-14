/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostVersionGetV10ResponseDataRaiseListInfosInner struct for QianchuanToolsSmartBoostAdBoostVersionGetV10ResponseDataRaiseListInfosInner
type QianchuanToolsSmartBoostAdBoostVersionGetV10ResponseDataRaiseListInfosInner struct {
	// 起量版本号
	AdRaiseVersion *int64 `json:"ad_raise_version,omitempty"`
	// 当前起量版本结束时间，格式：2021-03-31 17:00:00
	EndTime *string `json:"end_time,omitempty"`
	// 当前版本起量开始时间，格式：2021-03-31 16:00:00
	StartTime *string `json:"start_time,omitempty"`
}
