/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAssetType
type ProjectListV30DataListAssetType string

// List of project_list_v3.0_data_list_asset_type
const (
	AWEME_ProjectListV30DataListAssetType         ProjectListV30DataListAssetType = "AWEME"
	ENTERPRISE_ProjectListV30DataListAssetType    ProjectListV30DataListAssetType = "ENTERPRISE"
	MICRO_APP_ProjectListV30DataListAssetType     ProjectListV30DataListAssetType = "MICRO_APP"
	ORANGE_ProjectListV30DataListAssetType        ProjectListV30DataListAssetType = "ORANGE"
	THIRDPARTY_ProjectListV30DataListAssetType    ProjectListV30DataListAssetType = "THIRDPARTY"
	WECHAT_APPLET_ProjectListV30DataListAssetType ProjectListV30DataListAssetType = "WECHAT_APPLET"
)

// All allowed values of ProjectListV30DataListAssetType enum
var AllowedProjectListV30DataListAssetTypeEnumValues = []ProjectListV30DataListAssetType{
	"AWEME",
	"ENTERPRISE",
	"MICRO_APP",
	"ORANGE",
	"THIRDPARTY",
	"WECHAT_APPLET",
}

// NewProjectListV30DataListAssetTypeFromValue returns a pointer to a valid ProjectListV30DataListAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAssetTypeFromValue(v string) (*ProjectListV30DataListAssetType, error) {
	ev := ProjectListV30DataListAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAssetType: valid values are %v", v, AllowedProjectListV30DataListAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAssetType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_asset_type value
func (v ProjectListV30DataListAssetType) Ptr() *ProjectListV30DataListAssetType {
	return &v
}
