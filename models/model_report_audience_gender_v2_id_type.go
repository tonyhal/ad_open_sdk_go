/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceGenderV2IdType
type ReportAudienceGenderV2IdType string

// List of report_audience_gender_v2_id_type
const (
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN_ReportAudienceGenderV2IdType   ReportAudienceGenderV2IdType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
	AUDIENCE_STAT_ID_TYPE_ADVERTISER_ReportAudienceGenderV2IdType ReportAudienceGenderV2IdType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	AUDIENCE_STAT_ID_TYPE_AD_ReportAudienceGenderV2IdType         ReportAudienceGenderV2IdType = "AUDIENCE_STAT_ID_TYPE_AD"
)

// Ptr returns reference to report_audience_gender_v2_id_type value
func (v ReportAudienceGenderV2IdType) Ptr() *ReportAudienceGenderV2IdType {
	return &v
}
