/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserFundGetV2GrantTypeSplit
type AdvertiserFundGetV2GrantTypeSplit string

// List of advertiser_fund_get_v2_grant_type_split
const (
	ON_AdvertiserFundGetV2GrantTypeSplit  AdvertiserFundGetV2GrantTypeSplit = "ON"
	OFF_AdvertiserFundGetV2GrantTypeSplit AdvertiserFundGetV2GrantTypeSplit = "OFF"
)

// All allowed values of AdvertiserFundGetV2GrantTypeSplit enum
var AllowedAdvertiserFundGetV2GrantTypeSplitEnumValues = []AdvertiserFundGetV2GrantTypeSplit{
	"ON",
	"OFF",
}

// NewAdvertiserFundGetV2GrantTypeSplitFromValue returns a pointer to a valid AdvertiserFundGetV2GrantTypeSplit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserFundGetV2GrantTypeSplitFromValue(v string) (*AdvertiserFundGetV2GrantTypeSplit, error) {
	ev := AdvertiserFundGetV2GrantTypeSplit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserFundGetV2GrantTypeSplit: valid values are %v", v, AllowedAdvertiserFundGetV2GrantTypeSplitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserFundGetV2GrantTypeSplit) IsValid() bool {
	for _, existing := range AllowedAdvertiserFundGetV2GrantTypeSplitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_fund_get_v2_grant_type_split value
func (v AdvertiserFundGetV2GrantTypeSplit) Ptr() *AdvertiserFundGetV2GrantTypeSplit {
	return &v
}
