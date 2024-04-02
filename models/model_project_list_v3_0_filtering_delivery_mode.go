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

// ProjectListV30FilteringDeliveryMode
type ProjectListV30FilteringDeliveryMode string

// List of project_list_v3.0_filtering_delivery_mode
const (
	MANUAL_ProjectListV30FilteringDeliveryMode     ProjectListV30FilteringDeliveryMode = "MANUAL"
	PROCEDURAL_ProjectListV30FilteringDeliveryMode ProjectListV30FilteringDeliveryMode = "PROCEDURAL"
)

// All allowed values of ProjectListV30FilteringDeliveryMode enum
var AllowedProjectListV30FilteringDeliveryModeEnumValues = []ProjectListV30FilteringDeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewProjectListV30FilteringDeliveryModeFromValue returns a pointer to a valid ProjectListV30FilteringDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringDeliveryModeFromValue(v string) (*ProjectListV30FilteringDeliveryMode, error) {
	ev := ProjectListV30FilteringDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringDeliveryMode: valid values are %v", v, AllowedProjectListV30FilteringDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringDeliveryMode) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_delivery_mode value
func (v ProjectListV30FilteringDeliveryMode) Ptr() *ProjectListV30FilteringDeliveryMode {
	return &v
}
