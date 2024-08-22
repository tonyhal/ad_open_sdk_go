/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30KeywordsMatchType
type ProjectCreateV30KeywordsMatchType string

// List of project_create_v3.0_keywords_match_type
const (
	EXTENSIVE_ProjectCreateV30KeywordsMatchType ProjectCreateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_ProjectCreateV30KeywordsMatchType    ProjectCreateV30KeywordsMatchType = "PHRASE"
	PRECISION_ProjectCreateV30KeywordsMatchType ProjectCreateV30KeywordsMatchType = "PRECISION"
)

// Ptr returns reference to project_create_v3.0_keywords_match_type value
func (v ProjectCreateV30KeywordsMatchType) Ptr() *ProjectCreateV30KeywordsMatchType {
	return &v
}
