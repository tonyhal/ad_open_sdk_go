/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringBlueFlowPackageSetting
type ProjectListV30FilteringBlueFlowPackageSetting string

// List of project_list_v3.0_filtering_blue_flow_package_setting
const (
	OFF_ProjectListV30FilteringBlueFlowPackageSetting ProjectListV30FilteringBlueFlowPackageSetting = "OFF"
	ON_ProjectListV30FilteringBlueFlowPackageSetting  ProjectListV30FilteringBlueFlowPackageSetting = "ON"
)

// All allowed values of ProjectListV30FilteringBlueFlowPackageSetting enum
var AllowedProjectListV30FilteringBlueFlowPackageSettingEnumValues = []ProjectListV30FilteringBlueFlowPackageSetting{
	"OFF",
	"ON",
}

// NewProjectListV30FilteringBlueFlowPackageSettingFromValue returns a pointer to a valid ProjectListV30FilteringBlueFlowPackageSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringBlueFlowPackageSettingFromValue(v string) (*ProjectListV30FilteringBlueFlowPackageSetting, error) {
	ev := ProjectListV30FilteringBlueFlowPackageSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringBlueFlowPackageSetting: valid values are %v", v, AllowedProjectListV30FilteringBlueFlowPackageSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringBlueFlowPackageSetting) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringBlueFlowPackageSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_blue_flow_package_setting value
func (v ProjectListV30FilteringBlueFlowPackageSetting) Ptr() *ProjectListV30FilteringBlueFlowPackageSetting {
	return &v
}
