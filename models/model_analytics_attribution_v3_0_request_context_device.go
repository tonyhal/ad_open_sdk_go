/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AnalyticsAttributionV30RequestContextDevice
type AnalyticsAttributionV30RequestContextDevice struct {
	// 下单用户的模糊手机号，目前支持以下3种类型： 1. （新）仅后四位：例如*******1234，前七位需要用星号表示；当上传此手机号格式时，receiver_province、receiver_city必填，否则无法上报和正确归因 2. 省略中间四位：例如130****1234，中间四位需用星号表示 3. 原始手机号sha256后的结果，64位字符串 【注意】手机号的加密步骤仅在能获取原值手机号情况下，使用sha256加密，其他两种手机号形式切勿加密！否则会导致归因为0
	PhoneNumBlurred *string `json:"phone_num_blurred,omitempty"`
}
