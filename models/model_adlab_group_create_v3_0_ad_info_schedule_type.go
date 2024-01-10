/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoScheduleType
type AdlabGroupCreateV30AdInfoScheduleType string

// List of adlab_group_create_v3.0_ad_info_schedule_type
const (
	SCHEDULE_FROM_NOW_AdlabGroupCreateV30AdInfoScheduleType  AdlabGroupCreateV30AdInfoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_AdlabGroupCreateV30AdInfoScheduleType AdlabGroupCreateV30AdInfoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of AdlabGroupCreateV30AdInfoScheduleType enum
var AllowedAdlabGroupCreateV30AdInfoScheduleTypeEnumValues = []AdlabGroupCreateV30AdInfoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewAdlabGroupCreateV30AdInfoScheduleTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoScheduleTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoScheduleType, error) {
	ev := AdlabGroupCreateV30AdInfoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoScheduleType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoScheduleType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_schedule_type value
func (v AdlabGroupCreateV30AdInfoScheduleType) Ptr() *AdlabGroupCreateV30AdInfoScheduleType {
	return &v
}
