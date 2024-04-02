/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoAdTarget
type AdlabGroupCreateV30AdInfoAdTarget string

// List of adlab_group_create_v3.0_ad_info_ad_target
const (
	SMART_BID_CONSERVATIVE_AdlabGroupCreateV30AdInfoAdTarget AdlabGroupCreateV30AdInfoAdTarget = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_AdlabGroupCreateV30AdInfoAdTarget       AdlabGroupCreateV30AdInfoAdTarget = "SMART_BID_CUSTOM"
)

// All allowed values of AdlabGroupCreateV30AdInfoAdTarget enum
var AllowedAdlabGroupCreateV30AdInfoAdTargetEnumValues = []AdlabGroupCreateV30AdInfoAdTarget{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewAdlabGroupCreateV30AdInfoAdTargetFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAdTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAdTargetFromValue(v string) (*AdlabGroupCreateV30AdInfoAdTarget, error) {
	ev := AdlabGroupCreateV30AdInfoAdTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAdTarget: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAdTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAdTarget) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAdTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_ad_target value
func (v AdlabGroupCreateV30AdInfoAdTarget) Ptr() *AdlabGroupCreateV30AdInfoAdTarget {
	return &v
}
