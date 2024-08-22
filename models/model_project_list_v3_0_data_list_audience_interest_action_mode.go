/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30DataListAudienceInterestActionMode
type ProjectListV30DataListAudienceInterestActionMode string

// List of project_list_v3.0_data_list_audience_interest_action_mode
const (
	CUSTOM_ProjectListV30DataListAudienceInterestActionMode    ProjectListV30DataListAudienceInterestActionMode = "CUSTOM"
	RECOMMEND_ProjectListV30DataListAudienceInterestActionMode ProjectListV30DataListAudienceInterestActionMode = "RECOMMEND"
	UNLIMITED_ProjectListV30DataListAudienceInterestActionMode ProjectListV30DataListAudienceInterestActionMode = "UNLIMITED"
)

// Ptr returns reference to project_list_v3.0_data_list_audience_interest_action_mode value
func (v ProjectListV30DataListAudienceInterestActionMode) Ptr() *ProjectListV30DataListAudienceInterestActionMode {
	return &v
}
