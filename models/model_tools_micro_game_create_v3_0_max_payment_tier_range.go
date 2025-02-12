/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameCreateV30MaxPaymentTierRange
type ToolsMicroGameCreateV30MaxPaymentTierRange string

// List of tools_micro_game_create_v3.0_max_payment_tier_range
const (
	ABOVE_1000_ToolsMicroGameCreateV30MaxPaymentTierRange       ToolsMicroGameCreateV30MaxPaymentTierRange = "ABOVE_1000"
	BELOW_500_ToolsMicroGameCreateV30MaxPaymentTierRange        ToolsMicroGameCreateV30MaxPaymentTierRange = "BELOW_500"
	FROM_500_TO_1000_ToolsMicroGameCreateV30MaxPaymentTierRange ToolsMicroGameCreateV30MaxPaymentTierRange = "FROM_500_TO_1000"
)

// Ptr returns reference to tools_micro_game_create_v3.0_max_payment_tier_range value
func (v ToolsMicroGameCreateV30MaxPaymentTierRange) Ptr() *ToolsMicroGameCreateV30MaxPaymentTierRange {
	return &v
}
