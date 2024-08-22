/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStrategyListV2StrategyTypes
type CreativeStrategyListV2StrategyTypes string

// List of creative_strategy_list_v2_strategy_types
const (
	HORIZONTAL2_HORIZONTAL_CreativeStrategyListV2StrategyTypes CreativeStrategyListV2StrategyTypes = "Horizontal2Horizontal"
	HORIZONTAL2_VERTICAL_CreativeStrategyListV2StrategyTypes   CreativeStrategyListV2StrategyTypes = "Horizontal2Vertical"
	VERTICAL2_HORIZONTAL_CreativeStrategyListV2StrategyTypes   CreativeStrategyListV2StrategyTypes = "Vertical2Horizontal"
	VERTICAL2_VERTICAL_CreativeStrategyListV2StrategyTypes     CreativeStrategyListV2StrategyTypes = "Vertical2Vertical"
)

// Ptr returns reference to creative_strategy_list_v2_strategy_types value
func (v CreativeStrategyListV2StrategyTypes) Ptr() *CreativeStrategyListV2StrategyTypes {
	return &v
}
