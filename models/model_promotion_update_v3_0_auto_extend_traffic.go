/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30AutoExtendTraffic
type PromotionUpdateV30AutoExtendTraffic string

// List of promotion_update_v3.0_auto_extend_traffic
const (
	OFF_PromotionUpdateV30AutoExtendTraffic PromotionUpdateV30AutoExtendTraffic = "OFF"
	ON_PromotionUpdateV30AutoExtendTraffic  PromotionUpdateV30AutoExtendTraffic = "ON"
)

// All allowed values of PromotionUpdateV30AutoExtendTraffic enum
var AllowedPromotionUpdateV30AutoExtendTrafficEnumValues = []PromotionUpdateV30AutoExtendTraffic{
	"OFF",
	"ON",
}

// NewPromotionUpdateV30AutoExtendTrafficFromValue returns a pointer to a valid PromotionUpdateV30AutoExtendTraffic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30AutoExtendTrafficFromValue(v string) (*PromotionUpdateV30AutoExtendTraffic, error) {
	ev := PromotionUpdateV30AutoExtendTraffic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30AutoExtendTraffic: valid values are %v", v, AllowedPromotionUpdateV30AutoExtendTrafficEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30AutoExtendTraffic) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30AutoExtendTrafficEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_auto_extend_traffic value
func (v PromotionUpdateV30AutoExtendTraffic) Ptr() *PromotionUpdateV30AutoExtendTraffic {
	return &v
}
