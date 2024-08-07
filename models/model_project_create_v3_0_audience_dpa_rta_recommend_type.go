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

// ProjectCreateV30AudienceDpaRtaRecommendType
type ProjectCreateV30AudienceDpaRtaRecommendType string

// List of project_create_v3.0_audience_dpa_rta_recommend_type
const (
	MORE_ProjectCreateV30AudienceDpaRtaRecommendType ProjectCreateV30AudienceDpaRtaRecommendType = "MORE"
	ONLY_ProjectCreateV30AudienceDpaRtaRecommendType ProjectCreateV30AudienceDpaRtaRecommendType = "ONLY"
)

// All allowed values of ProjectCreateV30AudienceDpaRtaRecommendType enum
var AllowedProjectCreateV30AudienceDpaRtaRecommendTypeEnumValues = []ProjectCreateV30AudienceDpaRtaRecommendType{
	"MORE",
	"ONLY",
}

// NewProjectCreateV30AudienceDpaRtaRecommendTypeFromValue returns a pointer to a valid ProjectCreateV30AudienceDpaRtaRecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDpaRtaRecommendTypeFromValue(v string) (*ProjectCreateV30AudienceDpaRtaRecommendType, error) {
	ev := ProjectCreateV30AudienceDpaRtaRecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDpaRtaRecommendType: valid values are %v", v, AllowedProjectCreateV30AudienceDpaRtaRecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDpaRtaRecommendType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDpaRtaRecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_dpa_rta_recommend_type value
func (v ProjectCreateV30AudienceDpaRtaRecommendType) Ptr() *ProjectCreateV30AudienceDpaRtaRecommendType {
	return &v
}
