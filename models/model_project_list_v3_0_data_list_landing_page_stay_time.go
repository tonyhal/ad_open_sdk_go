/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListLandingPageStayTime
type ProjectListV30DataListLandingPageStayTime int64

// List of project_list_v3.0_data_list_landing_page_stay_time
const (
	Enum_1000_ProjectListV30DataListLandingPageStayTime  ProjectListV30DataListLandingPageStayTime = 1000
	Enum_12000_ProjectListV30DataListLandingPageStayTime ProjectListV30DataListLandingPageStayTime = 12000
	Enum_20000_ProjectListV30DataListLandingPageStayTime ProjectListV30DataListLandingPageStayTime = 20000
	Enum_30000_ProjectListV30DataListLandingPageStayTime ProjectListV30DataListLandingPageStayTime = 30000
	Enum_40000_ProjectListV30DataListLandingPageStayTime ProjectListV30DataListLandingPageStayTime = 40000
	Enum_5000_ProjectListV30DataListLandingPageStayTime  ProjectListV30DataListLandingPageStayTime = 5000
	Enum_50000_ProjectListV30DataListLandingPageStayTime ProjectListV30DataListLandingPageStayTime = 50000
	Enum_60000_ProjectListV30DataListLandingPageStayTime ProjectListV30DataListLandingPageStayTime = 60000
)

// All allowed values of ProjectListV30DataListLandingPageStayTime enum
var AllowedProjectListV30DataListLandingPageStayTimeEnumValues = []ProjectListV30DataListLandingPageStayTime{
	1000,
	12000,
	20000,
	30000,
	40000,
	5000,
	50000,
	60000,
}

// NewProjectListV30DataListLandingPageStayTimeFromValue returns a pointer to a valid ProjectListV30DataListLandingPageStayTime
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListLandingPageStayTimeFromValue(v int64) (*ProjectListV30DataListLandingPageStayTime, error) {
	ev := ProjectListV30DataListLandingPageStayTime(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListLandingPageStayTime: valid values are %v", v, AllowedProjectListV30DataListLandingPageStayTimeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListLandingPageStayTime) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListLandingPageStayTimeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_landing_page_stay_time value
func (v ProjectListV30DataListLandingPageStayTime) Ptr() *ProjectListV30DataListLandingPageStayTime {
	return &v
}
