/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigV2CreateV2RequestStrategyDataInner struct for CreativeAutoGenerateConfigV2CreateV2RequestStrategyDataInner
type CreativeAutoGenerateConfigV2CreateV2RequestStrategyDataInner struct {
	// 策略id
	StrategyId int64 `json:"strategy_id"`
	// 策略配置项列表
	StrategyState []*CreativeAutoGenerateConfigV2CreateV2RequestStrategyDataInnerStrategyStateInner `json:"strategy_state,omitempty"`
}
