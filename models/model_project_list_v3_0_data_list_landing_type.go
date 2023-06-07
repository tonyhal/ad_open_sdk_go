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

// ProjectListV30DataListLandingType
type ProjectListV30DataListLandingType string

// List of project_list_v3.0_data_list_landing_type
const (
	APP_ProjectListV30DataListLandingType            ProjectListV30DataListLandingType = "APP"
	ARTICLE_ProjectListV30DataListLandingType        ProjectListV30DataListLandingType = "ARTICLE"
	BRAND_EXTERNAL_ProjectListV30DataListLandingType ProjectListV30DataListLandingType = "BRAND_EXTERNAL"
	DPA_ProjectListV30DataListLandingType            ProjectListV30DataListLandingType = "DPA"
	GOODS_ProjectListV30DataListLandingType          ProjectListV30DataListLandingType = "GOODS"
	LINK_ProjectListV30DataListLandingType           ProjectListV30DataListLandingType = "LINK"
	LIVE_ProjectListV30DataListLandingType           ProjectListV30DataListLandingType = "LIVE"
	MICRO_GAME_ProjectListV30DataListLandingType     ProjectListV30DataListLandingType = "MICRO_GAME"
	NATIVE_ACTION_ProjectListV30DataListLandingType  ProjectListV30DataListLandingType = "NATIVE_ACTION"
	QUICK_APP_ProjectListV30DataListLandingType      ProjectListV30DataListLandingType = "QUICK_APP"
	SHOP_ProjectListV30DataListLandingType           ProjectListV30DataListLandingType = "SHOP"
	STORE_ProjectListV30DataListLandingType          ProjectListV30DataListLandingType = "STORE"
)

// All allowed values of ProjectListV30DataListLandingType enum
var AllowedProjectListV30DataListLandingTypeEnumValues = []ProjectListV30DataListLandingType{
	"APP",
	"ARTICLE",
	"BRAND_EXTERNAL",
	"DPA",
	"GOODS",
	"LINK",
	"LIVE",
	"MICRO_GAME",
	"NATIVE_ACTION",
	"QUICK_APP",
	"SHOP",
	"STORE",
}

// NewProjectListV30DataListLandingTypeFromValue returns a pointer to a valid ProjectListV30DataListLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListLandingTypeFromValue(v string) (*ProjectListV30DataListLandingType, error) {
	ev := ProjectListV30DataListLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListLandingType: valid values are %v", v, AllowedProjectListV30DataListLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListLandingType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_landing_type value
func (v ProjectListV30DataListLandingType) Ptr() *ProjectListV30DataListLandingType {
	return &v
}
