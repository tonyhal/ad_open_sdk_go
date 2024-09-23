/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectCreateV30AudienceHideIfConverted
type LocalProjectCreateV30AudienceHideIfConverted string

// List of local_project_create_v3.0_audience_hide_if_converted
const (
	ADVERTISER_LocalProjectCreateV30AudienceHideIfConverted   LocalProjectCreateV30AudienceHideIfConverted = "ADVERTISER"
	CUSTOMER_LocalProjectCreateV30AudienceHideIfConverted     LocalProjectCreateV30AudienceHideIfConverted = "CUSTOMER"
	ORGANIZATION_LocalProjectCreateV30AudienceHideIfConverted LocalProjectCreateV30AudienceHideIfConverted = "ORGANIZATION"
	PROJECT_LocalProjectCreateV30AudienceHideIfConverted      LocalProjectCreateV30AudienceHideIfConverted = "PROJECT"
	PROMOTION_LocalProjectCreateV30AudienceHideIfConverted    LocalProjectCreateV30AudienceHideIfConverted = "PROMOTION"
)

// Ptr returns reference to local_project_create_v3.0_audience_hide_if_converted value
func (v LocalProjectCreateV30AudienceHideIfConverted) Ptr() *LocalProjectCreateV30AudienceHideIfConverted {
	return &v
}
