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

// ProjectListV30DataListAudienceActionScene
type ProjectListV30DataListAudienceActionScene string

// List of project_list_v3.0_data_list_audience_action_scene
const (
	AD_ProjectListV30DataListAudienceActionScene         ProjectListV30DataListAudienceActionScene = "AD"
	APP_ProjectListV30DataListAudienceActionScene        ProjectListV30DataListAudienceActionScene = "APP"
	E_COMMERCE_ProjectListV30DataListAudienceActionScene ProjectListV30DataListAudienceActionScene = "E-COMMERCE"
	NEWS_ProjectListV30DataListAudienceActionScene       ProjectListV30DataListAudienceActionScene = "NEWS"
	SEARCH_ProjectListV30DataListAudienceActionScene     ProjectListV30DataListAudienceActionScene = "SEARCH"
)

// All allowed values of ProjectListV30DataListAudienceActionScene enum
var AllowedProjectListV30DataListAudienceActionSceneEnumValues = []ProjectListV30DataListAudienceActionScene{
	"AD",
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewProjectListV30DataListAudienceActionSceneFromValue returns a pointer to a valid ProjectListV30DataListAudienceActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceActionSceneFromValue(v string) (*ProjectListV30DataListAudienceActionScene, error) {
	ev := ProjectListV30DataListAudienceActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceActionScene: valid values are %v", v, AllowedProjectListV30DataListAudienceActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceActionScene) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_action_scene value
func (v ProjectListV30DataListAudienceActionScene) Ptr() *ProjectListV30DataListAudienceActionScene {
	return &v
}
