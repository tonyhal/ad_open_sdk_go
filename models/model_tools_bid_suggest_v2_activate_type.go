/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2ActivateType
type ToolsBidSuggestV2ActivateType string

// List of tools_bid_suggest_v2_activate_type
const (
	WITH_IN_A_MONTH_ToolsBidSuggestV2ActivateType         ToolsBidSuggestV2ActivateType = "WITH_IN_A_MONTH"
	THREE_MONTH_EAILIER_ToolsBidSuggestV2ActivateType     ToolsBidSuggestV2ActivateType = "THREE_MONTH_EAILIER"
	UNLIMITED_ToolsBidSuggestV2ActivateType               ToolsBidSuggestV2ActivateType = "UNLIMITED"
	ONE_MONTH_2_THREE_MONTH_ToolsBidSuggestV2ActivateType ToolsBidSuggestV2ActivateType = "ONE_MONTH_2_THREE_MONTH"
)

// Ptr returns reference to tools_bid_suggest_v2_activate_type value
func (v ToolsBidSuggestV2ActivateType) Ptr() *ToolsBidSuggestV2ActivateType {
	return &v
}
