/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AppPromotionType
type ProjectCreateV30AppPromotionType string

// List of project_create_v3.0_app_promotion_type
const (
	DOWNLOAD_ProjectCreateV30AppPromotionType ProjectCreateV30AppPromotionType = "DOWNLOAD"
	LAUNCH_ProjectCreateV30AppPromotionType   ProjectCreateV30AppPromotionType = "LAUNCH"
	RESERVE_ProjectCreateV30AppPromotionType  ProjectCreateV30AppPromotionType = "RESERVE"
)

// All allowed values of ProjectCreateV30AppPromotionType enum
var AllowedProjectCreateV30AppPromotionTypeEnumValues = []ProjectCreateV30AppPromotionType{
	"DOWNLOAD",
	"LAUNCH",
	"RESERVE",
}

// NewProjectCreateV30AppPromotionTypeFromValue returns a pointer to a valid ProjectCreateV30AppPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AppPromotionTypeFromValue(v string) (*ProjectCreateV30AppPromotionType, error) {
	ev := ProjectCreateV30AppPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AppPromotionType: valid values are %v", v, AllowedProjectCreateV30AppPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AppPromotionType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AppPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_app_promotion_type value
func (v ProjectCreateV30AppPromotionType) Ptr() *ProjectCreateV30AppPromotionType {
	return &v
}
