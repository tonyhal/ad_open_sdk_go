/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStrategyListV2ResponseData
type CreativeStrategyListV2ResponseData struct {
	PageInfo *CreativeStrategyListV2ResponseDataPageInfo `json:"page_info,omitempty"`
	// 策略列表
	StrategyModels []*CreativeStrategyListV2ResponseDataStrategyModelsInner `json:"strategy_models,omitempty"`
}
