/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasCreateBoostItemGroupV2ResponseData
type StarVasCreateBoostItemGroupV2ResponseData struct {
	// 下单失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// 是否下单成功
	Success bool `json:"success"`
}
