/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityScoreTotalGetV30DataScoreInfoListIllegalType
type SecurityScoreTotalGetV30DataScoreInfoListIllegalType string

// List of security_score_total_get_v3.0_data_score_info_list_illegal_type
const (
	GENERAL_SecurityScoreTotalGetV30DataScoreInfoListIllegalType       SecurityScoreTotalGetV30DataScoreInfoListIllegalType = "GENERAL"
	ONECLASS_SecurityScoreTotalGetV30DataScoreInfoListIllegalType      SecurityScoreTotalGetV30DataScoreInfoListIllegalType = "ONECLASS"
	SERIOUS_SecurityScoreTotalGetV30DataScoreInfoListIllegalType       SecurityScoreTotalGetV30DataScoreInfoListIllegalType = "SERIOUS"
	TWOTHREECLASS_SecurityScoreTotalGetV30DataScoreInfoListIllegalType SecurityScoreTotalGetV30DataScoreInfoListIllegalType = "TWOTHREECLASS"
)

// Ptr returns reference to security_score_total_get_v3.0_data_score_info_list_illegal_type value
func (v SecurityScoreTotalGetV30DataScoreInfoListIllegalType) Ptr() *SecurityScoreTotalGetV30DataScoreInfoListIllegalType {
	return &v
}
