/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderTerminateV10ResponseData
type QianchuanAwemeOrderTerminateV10ResponseData struct {
	// 终止成功的订单ID
	OrderIds           []int64                                                        `json:"order_ids,omitempty"`
	TerminateQuotaInfo *QianchuanAwemeOrderTerminateV10ResponseDataTerminateQuotaInfo `json:"terminate_quota_info,omitempty"`
}
