/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataPricing
type AdGetV2DataPricing string

// List of ad_get_v2_data_pricing
const (
	PRICING_CPA_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPA"
	PRICING_CPC_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPC"
	PRICING_CPV_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPV"
	PRICING_CPM_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPM"
	PRICING_CPC_OCPM_AdGetV2DataPricing AdGetV2DataPricing = "PRICING_CPC_OCPM"
	PRICING_OCPC_AdGetV2DataPricing     AdGetV2DataPricing = "PRICING_OCPC"
	PRICING_OCPM_AdGetV2DataPricing     AdGetV2DataPricing = "PRICING_OCPM"
)

// All allowed values of AdGetV2DataPricing enum
var AllowedAdGetV2DataPricingEnumValues = []AdGetV2DataPricing{
	"PRICING_CPA",
	"PRICING_CPC",
	"PRICING_CPV",
	"PRICING_CPM",
	"PRICING_CPC_OCPM",
	"PRICING_OCPC",
	"PRICING_OCPM",
}

// NewAdGetV2DataPricingFromValue returns a pointer to a valid AdGetV2DataPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataPricingFromValue(v string) (*AdGetV2DataPricing, error) {
	ev := AdGetV2DataPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataPricing: valid values are %v", v, AllowedAdGetV2DataPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataPricing) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_pricing value
func (v AdGetV2DataPricing) Ptr() *AdGetV2DataPricing {
	return &v
}