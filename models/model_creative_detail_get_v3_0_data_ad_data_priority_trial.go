/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataAdDataPriorityTrial
type CreativeDetailGetV30DataAdDataPriorityTrial string

// List of creative_detail_get_v3.0_data_ad_data_priority_trial
const (
	OFF_CreativeDetailGetV30DataAdDataPriorityTrial CreativeDetailGetV30DataAdDataPriorityTrial = "OFF"
	ON_CreativeDetailGetV30DataAdDataPriorityTrial  CreativeDetailGetV30DataAdDataPriorityTrial = "ON"
)

// All allowed values of CreativeDetailGetV30DataAdDataPriorityTrial enum
var AllowedCreativeDetailGetV30DataAdDataPriorityTrialEnumValues = []CreativeDetailGetV30DataAdDataPriorityTrial{
	"OFF",
	"ON",
}

// NewCreativeDetailGetV30DataAdDataPriorityTrialFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataPriorityTrial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataPriorityTrialFromValue(v string) (*CreativeDetailGetV30DataAdDataPriorityTrial, error) {
	ev := CreativeDetailGetV30DataAdDataPriorityTrial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataPriorityTrial: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataPriorityTrialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataPriorityTrial) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataPriorityTrialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_priority_trial value
func (v CreativeDetailGetV30DataAdDataPriorityTrial) Ptr() *CreativeDetailGetV30DataAdDataPriorityTrial {
	return &v
}
