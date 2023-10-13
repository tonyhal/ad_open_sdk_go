/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListRelatedProductProductSetting
type ProjectListV30DataListRelatedProductProductSetting string

// List of project_list_v3.0_data_list_related_product_product_setting
const (
	NO_MAP_ProjectListV30DataListRelatedProductProductSetting ProjectListV30DataListRelatedProductProductSetting = "NO_MAP"
	SINGLE_ProjectListV30DataListRelatedProductProductSetting ProjectListV30DataListRelatedProductProductSetting = "SINGLE"
)

// All allowed values of ProjectListV30DataListRelatedProductProductSetting enum
var AllowedProjectListV30DataListRelatedProductProductSettingEnumValues = []ProjectListV30DataListRelatedProductProductSetting{
	"NO_MAP",
	"SINGLE",
}

// NewProjectListV30DataListRelatedProductProductSettingFromValue returns a pointer to a valid ProjectListV30DataListRelatedProductProductSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListRelatedProductProductSettingFromValue(v string) (*ProjectListV30DataListRelatedProductProductSetting, error) {
	ev := ProjectListV30DataListRelatedProductProductSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListRelatedProductProductSetting: valid values are %v", v, AllowedProjectListV30DataListRelatedProductProductSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListRelatedProductProductSetting) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListRelatedProductProductSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_related_product_product_setting value
func (v ProjectListV30DataListRelatedProductProductSetting) Ptr() *ProjectListV30DataListRelatedProductProductSetting {
	return &v
}
