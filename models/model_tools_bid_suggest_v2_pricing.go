/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2Pricing
type ToolsBidSuggestV2Pricing string

// List of tools_bid_suggest_v2_pricing
const (
	PRICING_OCPM_ToolsBidSuggestV2Pricing     ToolsBidSuggestV2Pricing = "PRICING_OCPM"
	PRICING_CPC_OCPM_ToolsBidSuggestV2Pricing ToolsBidSuggestV2Pricing = "PRICING_CPC_OCPM"
	PRICING_CPV_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPV"
	PRICING_CPC_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPC"
	PRICING_OCPC_ToolsBidSuggestV2Pricing     ToolsBidSuggestV2Pricing = "PRICING_OCPC"
	PRICING_CPM_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPM"
	PRICING_CPA_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPA"
)

// All allowed values of ToolsBidSuggestV2Pricing enum
var AllowedToolsBidSuggestV2PricingEnumValues = []ToolsBidSuggestV2Pricing{
	"PRICING_OCPM",
	"PRICING_CPC_OCPM",
	"PRICING_CPV",
	"PRICING_CPC",
	"PRICING_OCPC",
	"PRICING_CPM",
	"PRICING_CPA",
}

// NewToolsBidSuggestV2PricingFromValue returns a pointer to a valid ToolsBidSuggestV2Pricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2PricingFromValue(v string) (*ToolsBidSuggestV2Pricing, error) {
	ev := ToolsBidSuggestV2Pricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Pricing: valid values are %v", v, AllowedToolsBidSuggestV2PricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Pricing) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2PricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_pricing value
func (v ToolsBidSuggestV2Pricing) Ptr() *ToolsBidSuggestV2Pricing {
	return &v
}
