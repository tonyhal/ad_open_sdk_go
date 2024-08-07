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

// ProjectListV30DataListAudienceAndroidOsv
type ProjectListV30DataListAudienceAndroidOsv string

// List of project_list_v3.0_data_list_audience_android_osv
const (
	Enum_10_0_ProjectListV30DataListAudienceAndroidOsv ProjectListV30DataListAudienceAndroidOsv = "10.0"
	Enum_10_1_ProjectListV30DataListAudienceAndroidOsv ProjectListV30DataListAudienceAndroidOsv = "10.1"
	Enum_11_0_ProjectListV30DataListAudienceAndroidOsv ProjectListV30DataListAudienceAndroidOsv = "11.0"
	Enum_2_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "2.0"
	Enum_2_1_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "2.1"
	Enum_2_2_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "2.2"
	Enum_2_3_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "2.3"
	Enum_3_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "3.0"
	Enum_3_1_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "3.1"
	Enum_3_2_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "3.2"
	Enum_4_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "4.0"
	Enum_4_1_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "4.1"
	Enum_4_2_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "4.2"
	Enum_4_3_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "4.3"
	Enum_4_4_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "4.4"
	Enum_4_5_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "4.5"
	Enum_5_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "5.0"
	Enum_5_1_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "5.1"
	Enum_6_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "6.0"
	Enum_7_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "7.0"
	Enum_7_1_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "7.1"
	Enum_8_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "8.0"
	Enum_8_1_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "8.1"
	Enum_9_0_ProjectListV30DataListAudienceAndroidOsv  ProjectListV30DataListAudienceAndroidOsv = "9.0"
	NONE_ProjectListV30DataListAudienceAndroidOsv      ProjectListV30DataListAudienceAndroidOsv = "NONE"
)

// All allowed values of ProjectListV30DataListAudienceAndroidOsv enum
var AllowedProjectListV30DataListAudienceAndroidOsvEnumValues = []ProjectListV30DataListAudienceAndroidOsv{
	"10.0",
	"10.1",
	"11.0",
	"2.0",
	"2.1",
	"2.2",
	"2.3",
	"3.0",
	"3.1",
	"3.2",
	"4.0",
	"4.1",
	"4.2",
	"4.3",
	"4.4",
	"4.5",
	"5.0",
	"5.1",
	"6.0",
	"7.0",
	"7.1",
	"8.0",
	"8.1",
	"9.0",
	"NONE",
}

// NewProjectListV30DataListAudienceAndroidOsvFromValue returns a pointer to a valid ProjectListV30DataListAudienceAndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceAndroidOsvFromValue(v string) (*ProjectListV30DataListAudienceAndroidOsv, error) {
	ev := ProjectListV30DataListAudienceAndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceAndroidOsv: valid values are %v", v, AllowedProjectListV30DataListAudienceAndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceAndroidOsv) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceAndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_android_osv value
func (v ProjectListV30DataListAudienceAndroidOsv) Ptr() *ProjectListV30DataListAudienceAndroidOsv {
	return &v
}
