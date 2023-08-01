/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingBidSpeed
type ProjectCreateV30DeliverySettingBidSpeed string

// List of project_create_v3.0_delivery_setting_bid_speed
const (
	BALANCE_ProjectCreateV30DeliverySettingBidSpeed ProjectCreateV30DeliverySettingBidSpeed = "BALANCE"
	FAST_ProjectCreateV30DeliverySettingBidSpeed    ProjectCreateV30DeliverySettingBidSpeed = "FAST"
)

// All allowed values of ProjectCreateV30DeliverySettingBidSpeed enum
var AllowedProjectCreateV30DeliverySettingBidSpeedEnumValues = []ProjectCreateV30DeliverySettingBidSpeed{
	"BALANCE",
	"FAST",
}

// NewProjectCreateV30DeliverySettingBidSpeedFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingBidSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingBidSpeedFromValue(v string) (*ProjectCreateV30DeliverySettingBidSpeed, error) {
	ev := ProjectCreateV30DeliverySettingBidSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingBidSpeed: valid values are %v", v, AllowedProjectCreateV30DeliverySettingBidSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingBidSpeed) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingBidSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_bid_speed value
func (v ProjectCreateV30DeliverySettingBidSpeed) Ptr() *ProjectCreateV30DeliverySettingBidSpeed {
	return &v
}