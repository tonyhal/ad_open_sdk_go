/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectUpdateV30AudienceHideIfConverted
type LocalProjectUpdateV30AudienceHideIfConverted string

// List of local_project_update_v3.0_audience_hide_if_converted
const (
	ADVERTISER_LocalProjectUpdateV30AudienceHideIfConverted   LocalProjectUpdateV30AudienceHideIfConverted = "ADVERTISER"
	CUSTOMER_LocalProjectUpdateV30AudienceHideIfConverted     LocalProjectUpdateV30AudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_LocalProjectUpdateV30AudienceHideIfConverted   LocalProjectUpdateV30AudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_LocalProjectUpdateV30AudienceHideIfConverted LocalProjectUpdateV30AudienceHideIfConverted = "ORGANIZATION"
	PROJECT_LocalProjectUpdateV30AudienceHideIfConverted      LocalProjectUpdateV30AudienceHideIfConverted = "PROJECT"
	PROMOTION_LocalProjectUpdateV30AudienceHideIfConverted    LocalProjectUpdateV30AudienceHideIfConverted = "PROMOTION"
)

// Ptr returns reference to local_project_update_v3.0_audience_hide_if_converted value
func (v LocalProjectUpdateV30AudienceHideIfConverted) Ptr() *LocalProjectUpdateV30AudienceHideIfConverted {
	return &v
}
