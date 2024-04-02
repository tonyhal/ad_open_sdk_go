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

// ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay
type ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay string

// List of tools_promotion_raise_set_v3.0_raise_info_appointed_time_repeated_day
const (
	EVERY_DAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay     ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "EVERY_DAY"
	PER_FRIDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay    ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_FRIDAY"
	PER_MONDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay    ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_MONDAY"
	PER_SATURDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay  ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_SATURDAY"
	PER_SUNDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay    ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_SUNDAY"
	PER_THURSDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay  ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_THURSDAY"
	PER_TUESDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay   ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_TUESDAY"
	PER_WEDNESDAY_ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay = "PER_WEDNESDAY"
)

// All allowed values of ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay enum
var AllowedToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDayEnumValues = []ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay{
	"EVERY_DAY",
	"PER_FRIDAY",
	"PER_MONDAY",
	"PER_SATURDAY",
	"PER_SUNDAY",
	"PER_THURSDAY",
	"PER_TUESDAY",
	"PER_WEDNESDAY",
}

// NewToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDayFromValue returns a pointer to a valid ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDayFromValue(v string) (*ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay, error) {
	ev := ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay: valid values are %v", v, AllowedToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay) IsValid() bool {
	for _, existing := range AllowedToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_raise_set_v3.0_raise_info_appointed_time_repeated_day value
func (v ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay) Ptr() *ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay {
	return &v
}
