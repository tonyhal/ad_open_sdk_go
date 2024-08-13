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

// ProjectListV30DataListMultiAssetType
type ProjectListV30DataListMultiAssetType string

// List of project_list_v3.0_data_list_multi_asset_type
const (
	ORANGE_AND_AWEME_ProjectListV30DataListMultiAssetType ProjectListV30DataListMultiAssetType = "ORANGE_AND_AWEME"
)

// All allowed values of ProjectListV30DataListMultiAssetType enum
var AllowedProjectListV30DataListMultiAssetTypeEnumValues = []ProjectListV30DataListMultiAssetType{
	"ORANGE_AND_AWEME",
}

// NewProjectListV30DataListMultiAssetTypeFromValue returns a pointer to a valid ProjectListV30DataListMultiAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListMultiAssetTypeFromValue(v string) (*ProjectListV30DataListMultiAssetType, error) {
	ev := ProjectListV30DataListMultiAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListMultiAssetType: valid values are %v", v, AllowedProjectListV30DataListMultiAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListMultiAssetType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListMultiAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_multi_asset_type value
func (v ProjectListV30DataListMultiAssetType) Ptr() *ProjectListV30DataListMultiAssetType {
	return &v
}
