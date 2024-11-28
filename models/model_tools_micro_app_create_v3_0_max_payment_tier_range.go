/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppCreateV30MaxPaymentTierRange
type ToolsMicroAppCreateV30MaxPaymentTierRange string

// List of tools_micro_app_create_v3.0_max_payment_tier_range
const (
	ABOVE_1000_ToolsMicroAppCreateV30MaxPaymentTierRange       ToolsMicroAppCreateV30MaxPaymentTierRange = "ABOVE_1000"
	BELOW_500_ToolsMicroAppCreateV30MaxPaymentTierRange        ToolsMicroAppCreateV30MaxPaymentTierRange = "BELOW_500"
	FROM_500_TO_1000_ToolsMicroAppCreateV30MaxPaymentTierRange ToolsMicroAppCreateV30MaxPaymentTierRange = "FROM_500_TO_1000"
)

// Ptr returns reference to tools_micro_app_create_v3.0_max_payment_tier_range value
func (v ToolsMicroAppCreateV30MaxPaymentTierRange) Ptr() *ToolsMicroAppCreateV30MaxPaymentTierRange {
	return &v
}