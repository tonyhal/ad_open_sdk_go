/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataDeliveryRange
type AdGetV2DataDeliveryRange string

// List of ad_get_v2_data_delivery_range
const (
	UNIVERSAL_AdGetV2DataDeliveryRange AdGetV2DataDeliveryRange = "UNIVERSAL"
	DEFAULT_AdGetV2DataDeliveryRange   AdGetV2DataDeliveryRange = "DEFAULT"
	UNION_AdGetV2DataDeliveryRange     AdGetV2DataDeliveryRange = "UNION"
)

// All allowed values of AdGetV2DataDeliveryRange enum
var AllowedAdGetV2DataDeliveryRangeEnumValues = []AdGetV2DataDeliveryRange{
	"UNIVERSAL",
	"DEFAULT",
	"UNION",
}

// NewAdGetV2DataDeliveryRangeFromValue returns a pointer to a valid AdGetV2DataDeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDeliveryRangeFromValue(v string) (*AdGetV2DataDeliveryRange, error) {
	ev := AdGetV2DataDeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDeliveryRange: valid values are %v", v, AllowedAdGetV2DataDeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDeliveryRange) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_delivery_range value
func (v AdGetV2DataDeliveryRange) Ptr() *AdGetV2DataDeliveryRange {
	return &v
}
