/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30AudienceHideIfExists
type ProjectCreateV30AudienceHideIfExists string

// List of project_create_v3.0_audience_hide_if_exists
const (
	FILTER_ProjectCreateV30AudienceHideIfExists    ProjectCreateV30AudienceHideIfExists = "FILTER"
	TARGETING_ProjectCreateV30AudienceHideIfExists ProjectCreateV30AudienceHideIfExists = "TARGETING"
	UNLIMITED_ProjectCreateV30AudienceHideIfExists ProjectCreateV30AudienceHideIfExists = "UNLIMITED"
)

// Ptr returns reference to project_create_v3.0_audience_hide_if_exists value
func (v ProjectCreateV30AudienceHideIfExists) Ptr() *ProjectCreateV30AudienceHideIfExists {
	return &v
}
