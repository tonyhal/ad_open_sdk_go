/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceAutoExtendTargets
type ProjectListV30DataListAudienceAutoExtendTargets string

// List of project_list_v3.0_data_list_audience_auto_extend_targets
const (
	AD_TAG_ProjectListV30DataListAudienceAutoExtendTargets          ProjectListV30DataListAudienceAutoExtendTargets = "AD_TAG"
	AGE_ProjectListV30DataListAudienceAutoExtendTargets             ProjectListV30DataListAudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_ProjectListV30DataListAudienceAutoExtendTargets ProjectListV30DataListAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_ProjectListV30DataListAudienceAutoExtendTargets          ProjectListV30DataListAudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_ProjectListV30DataListAudienceAutoExtendTargets ProjectListV30DataListAudienceAutoExtendTargets = "INTEREST_ACTION"
	INTEREST_TAG_ProjectListV30DataListAudienceAutoExtendTargets    ProjectListV30DataListAudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_ProjectListV30DataListAudienceAutoExtendTargets          ProjectListV30DataListAudienceAutoExtendTargets = "REGION"
)

// All allowed values of ProjectListV30DataListAudienceAutoExtendTargets enum
var AllowedProjectListV30DataListAudienceAutoExtendTargetsEnumValues = []ProjectListV30DataListAudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"INTEREST_TAG",
	"REGION",
}

// NewProjectListV30DataListAudienceAutoExtendTargetsFromValue returns a pointer to a valid ProjectListV30DataListAudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceAutoExtendTargetsFromValue(v string) (*ProjectListV30DataListAudienceAutoExtendTargets, error) {
	ev := ProjectListV30DataListAudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceAutoExtendTargets: valid values are %v", v, AllowedProjectListV30DataListAudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_auto_extend_targets value
func (v ProjectListV30DataListAudienceAutoExtendTargets) Ptr() *ProjectListV30DataListAudienceAutoExtendTargets {
	return &v
}
