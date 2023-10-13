/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderTerminateV10ResponseDataTerminateQuotaInfo 终止订单配额信息
type QianchuanAwemeOrderTerminateV10ResponseDataTerminateQuotaInfo struct {
	// 终止订单配额
	TerminateQuotaSum *int64 `json:"terminate_quota_sum,omitempty"`
	// 已终止订单数
	TerminateQuotaUsed *int64 `json:"terminate_quota_used,omitempty"`
}
