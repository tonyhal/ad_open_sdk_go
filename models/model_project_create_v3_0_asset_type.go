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

// ProjectCreateV30AssetType
type ProjectCreateV30AssetType string

// List of project_create_v3.0_asset_type
const (
	AWEME_ProjectCreateV30AssetType         ProjectCreateV30AssetType = "AWEME"
	ENTERPRISE_ProjectCreateV30AssetType    ProjectCreateV30AssetType = "ENTERPRISE"
	MICRO_APP_ProjectCreateV30AssetType     ProjectCreateV30AssetType = "MICRO_APP"
	ORANGE_ProjectCreateV30AssetType        ProjectCreateV30AssetType = "ORANGE"
	THIRDPARTY_ProjectCreateV30AssetType    ProjectCreateV30AssetType = "THIRDPARTY"
	WECHAT_APPLET_ProjectCreateV30AssetType ProjectCreateV30AssetType = "WECHAT_APPLET"
)

// All allowed values of ProjectCreateV30AssetType enum
var AllowedProjectCreateV30AssetTypeEnumValues = []ProjectCreateV30AssetType{
	"AWEME",
	"ENTERPRISE",
	"MICRO_APP",
	"ORANGE",
	"THIRDPARTY",
	"WECHAT_APPLET",
}

// NewProjectCreateV30AssetTypeFromValue returns a pointer to a valid ProjectCreateV30AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AssetTypeFromValue(v string) (*ProjectCreateV30AssetType, error) {
	ev := ProjectCreateV30AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AssetType: valid values are %v", v, AllowedProjectCreateV30AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AssetType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_asset_type value
func (v ProjectCreateV30AssetType) Ptr() *ProjectCreateV30AssetType {
	return &v
}
