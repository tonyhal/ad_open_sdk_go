/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliveryRangeUnionVideoType
type ProjectCreateV30DeliveryRangeUnionVideoType string

// List of project_create_v3.0_delivery_range_union_video_type
const (
	ORIGINAL_VIDEO_ProjectCreateV30DeliveryRangeUnionVideoType ProjectCreateV30DeliveryRangeUnionVideoType = "ORIGINAL_VIDEO"
	REWARDED_VIDEO_ProjectCreateV30DeliveryRangeUnionVideoType ProjectCreateV30DeliveryRangeUnionVideoType = "REWARDED_VIDEO"
	SPLASH_VIDEO_ProjectCreateV30DeliveryRangeUnionVideoType   ProjectCreateV30DeliveryRangeUnionVideoType = "SPLASH_VIDEO"
)

// All allowed values of ProjectCreateV30DeliveryRangeUnionVideoType enum
var AllowedProjectCreateV30DeliveryRangeUnionVideoTypeEnumValues = []ProjectCreateV30DeliveryRangeUnionVideoType{
	"ORIGINAL_VIDEO",
	"REWARDED_VIDEO",
	"SPLASH_VIDEO",
}

// NewProjectCreateV30DeliveryRangeUnionVideoTypeFromValue returns a pointer to a valid ProjectCreateV30DeliveryRangeUnionVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliveryRangeUnionVideoTypeFromValue(v string) (*ProjectCreateV30DeliveryRangeUnionVideoType, error) {
	ev := ProjectCreateV30DeliveryRangeUnionVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliveryRangeUnionVideoType: valid values are %v", v, AllowedProjectCreateV30DeliveryRangeUnionVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliveryRangeUnionVideoType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliveryRangeUnionVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_range_union_video_type value
func (v ProjectCreateV30DeliveryRangeUnionVideoType) Ptr() *ProjectCreateV30DeliveryRangeUnionVideoType {
	return &v
}
