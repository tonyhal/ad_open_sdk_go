/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionRejectReasonGetV30ResponseData
type PromotionRejectReasonGetV30ResponseData struct {
	// 广告审核建议列表
	List []*PromotionRejectReasonGetV30ResponseDataListInner `json:"list,omitempty"`
}