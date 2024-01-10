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

// ProjectListV30DataListDeliverySettingScheduleType
type ProjectListV30DataListDeliverySettingScheduleType string

// List of project_list_v3.0_data_list_delivery_setting_schedule_type
const (
	SCHEDULE_FROM_NOW_ProjectListV30DataListDeliverySettingScheduleType  ProjectListV30DataListDeliverySettingScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_ProjectListV30DataListDeliverySettingScheduleType ProjectListV30DataListDeliverySettingScheduleType = "SCHEDULE_START_END"
)

// All allowed values of ProjectListV30DataListDeliverySettingScheduleType enum
var AllowedProjectListV30DataListDeliverySettingScheduleTypeEnumValues = []ProjectListV30DataListDeliverySettingScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewProjectListV30DataListDeliverySettingScheduleTypeFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingScheduleTypeFromValue(v string) (*ProjectListV30DataListDeliverySettingScheduleType, error) {
	ev := ProjectListV30DataListDeliverySettingScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingScheduleType: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingScheduleType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_schedule_type value
func (v ProjectListV30DataListDeliverySettingScheduleType) Ptr() *ProjectListV30DataListDeliverySettingScheduleType {
	return &v
}
