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

// AdlabGroupListV30FilteringLandingType
type AdlabGroupListV30FilteringLandingType string

// List of adlab_group_list_v3.0_filtering_landing_type
const (
	APP_AdlabGroupListV30FilteringLandingType  AdlabGroupListV30FilteringLandingType = "APP"
	LINK_AdlabGroupListV30FilteringLandingType AdlabGroupListV30FilteringLandingType = "LINK"
)

// All allowed values of AdlabGroupListV30FilteringLandingType enum
var AllowedAdlabGroupListV30FilteringLandingTypeEnumValues = []AdlabGroupListV30FilteringLandingType{
	"APP",
	"LINK",
}

// NewAdlabGroupListV30FilteringLandingTypeFromValue returns a pointer to a valid AdlabGroupListV30FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30FilteringLandingTypeFromValue(v string) (*AdlabGroupListV30FilteringLandingType, error) {
	ev := AdlabGroupListV30FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30FilteringLandingType: valid values are %v", v, AllowedAdlabGroupListV30FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_filtering_landing_type value
func (v AdlabGroupListV30FilteringLandingType) Ptr() *AdlabGroupListV30FilteringLandingType {
	return &v
}
