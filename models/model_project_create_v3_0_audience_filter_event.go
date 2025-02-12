/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30AudienceFilterEvent
type ProjectCreateV30AudienceFilterEvent string

// List of project_create_v3.0_audience_filter_event
const (
	AD_CONVERT_EXTERNAL_ACTION_ProjectCreateV30AudienceFilterEvent      ProjectCreateV30AudienceFilterEvent = "AD_CONVERT_EXTERNAL_ACTION"
	AD_CONVERT_TYPE_ACTIVE_ProjectCreateV30AudienceFilterEvent          ProjectCreateV30AudienceFilterEvent = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_ProjectCreateV30AudienceFilterEvent ProjectCreateV30AudienceFilterEvent = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_PAY_ProjectCreateV30AudienceFilterEvent             ProjectCreateV30AudienceFilterEvent = "AD_CONVERT_TYPE_PAY"
)

// Ptr returns reference to project_create_v3.0_audience_filter_event value
func (v ProjectCreateV30AudienceFilterEvent) Ptr() *ProjectCreateV30AudienceFilterEvent {
	return &v
}
