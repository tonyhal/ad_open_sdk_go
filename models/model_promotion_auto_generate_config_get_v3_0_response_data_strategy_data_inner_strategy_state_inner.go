/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInnerStrategyStateInner struct for PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInnerStrategyStateInner
type PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInnerStrategyStateInner struct {
	// 配置项key
	StateKey  string                                                                  `json:"state_key"`
	StateType PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType `json:"state_type"`
	// 配置项值
	StateValue *string `json:"state_value,omitempty"`
}
