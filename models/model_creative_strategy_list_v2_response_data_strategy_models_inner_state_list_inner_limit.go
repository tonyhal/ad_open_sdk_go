/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInnerLimit 配置项限制条件
type CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInnerLimit struct {
	// 图片高度(仅对state_type=Image有效)
	ImgHeight *int64 `json:"img_height,omitempty"`
	// 图片宽度(仅对state_type=Image有效)
	ImgWidth *int64 `json:"img_width,omitempty"`
	// 文案最大长度(仅对state_type=Text有效)
	TextMaxLength *int64 `json:"text_max_length,omitempty"`
	// 文案最小长度(仅对state_type=Text有效)
	TextMinLength *int64 `json:"text_min_length,omitempty"`
}
