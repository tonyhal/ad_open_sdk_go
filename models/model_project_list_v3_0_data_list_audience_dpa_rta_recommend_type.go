/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceDpaRtaRecommendType
type ProjectListV30DataListAudienceDpaRtaRecommendType string

// List of project_list_v3.0_data_list_audience_dpa_rta_recommend_type
const (
	MORE_ProjectListV30DataListAudienceDpaRtaRecommendType ProjectListV30DataListAudienceDpaRtaRecommendType = "MORE"
	ONLY_ProjectListV30DataListAudienceDpaRtaRecommendType ProjectListV30DataListAudienceDpaRtaRecommendType = "ONLY"
)

// All allowed values of ProjectListV30DataListAudienceDpaRtaRecommendType enum
var AllowedProjectListV30DataListAudienceDpaRtaRecommendTypeEnumValues = []ProjectListV30DataListAudienceDpaRtaRecommendType{
	"MORE",
	"ONLY",
}

// NewProjectListV30DataListAudienceDpaRtaRecommendTypeFromValue returns a pointer to a valid ProjectListV30DataListAudienceDpaRtaRecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDpaRtaRecommendTypeFromValue(v string) (*ProjectListV30DataListAudienceDpaRtaRecommendType, error) {
	ev := ProjectListV30DataListAudienceDpaRtaRecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDpaRtaRecommendType: valid values are %v", v, AllowedProjectListV30DataListAudienceDpaRtaRecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDpaRtaRecommendType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDpaRtaRecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_dpa_rta_recommend_type value
func (v ProjectListV30DataListAudienceDpaRtaRecommendType) Ptr() *ProjectListV30DataListAudienceDpaRtaRecommendType {
	return &v
}
