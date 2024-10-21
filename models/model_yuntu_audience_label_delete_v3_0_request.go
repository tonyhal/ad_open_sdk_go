/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// YuntuAudienceLabelDeleteV30Request struct for YuntuAudienceLabelDeleteV30Request
type YuntuAudienceLabelDeleteV30Request struct {
	// 人群标签ID，需传入当前用户通过 API 渠道创建的标签ID（创建成功后会返回），否则会校验失败。
	LabelId int64 `json:"label_id"`
	// 服务商ID
	ServiceProviderId int64 `json:"service_provider_id"`
	// 品牌ID
	YuntuBrandId int64 `json:"yuntu_brand_id"`
}