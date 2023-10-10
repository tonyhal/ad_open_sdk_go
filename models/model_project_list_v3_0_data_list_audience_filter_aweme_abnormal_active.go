/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceFilterAwemeAbnormalActive
type ProjectListV30DataListAudienceFilterAwemeAbnormalActive string

// List of project_list_v3.0_data_list_audience_filter_aweme_abnormal_active
const (
	OFF_ProjectListV30DataListAudienceFilterAwemeAbnormalActive ProjectListV30DataListAudienceFilterAwemeAbnormalActive = "OFF"
	ON_ProjectListV30DataListAudienceFilterAwemeAbnormalActive  ProjectListV30DataListAudienceFilterAwemeAbnormalActive = "ON"
)

// All allowed values of ProjectListV30DataListAudienceFilterAwemeAbnormalActive enum
var AllowedProjectListV30DataListAudienceFilterAwemeAbnormalActiveEnumValues = []ProjectListV30DataListAudienceFilterAwemeAbnormalActive{
	"OFF",
	"ON",
}

// NewProjectListV30DataListAudienceFilterAwemeAbnormalActiveFromValue returns a pointer to a valid ProjectListV30DataListAudienceFilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceFilterAwemeAbnormalActiveFromValue(v string) (*ProjectListV30DataListAudienceFilterAwemeAbnormalActive, error) {
	ev := ProjectListV30DataListAudienceFilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceFilterAwemeAbnormalActive: valid values are %v", v, AllowedProjectListV30DataListAudienceFilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceFilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceFilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_filter_aweme_abnormal_active value
func (v ProjectListV30DataListAudienceFilterAwemeAbnormalActive) Ptr() *ProjectListV30DataListAudienceFilterAwemeAbnormalActive {
	return &v
}
