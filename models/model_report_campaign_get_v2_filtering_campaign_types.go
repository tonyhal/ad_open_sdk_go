/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2FilteringCampaignTypes
type ReportCampaignGetV2FilteringCampaignTypes string

// List of report_campaign_get_v2_filtering_campaign_types
const (
	FEED_ReportCampaignGetV2FilteringCampaignTypes    ReportCampaignGetV2FilteringCampaignTypes = "FEED"
	CONTENT_ReportCampaignGetV2FilteringCampaignTypes ReportCampaignGetV2FilteringCampaignTypes = "CONTENT"
	SEARCH_ReportCampaignGetV2FilteringCampaignTypes  ReportCampaignGetV2FilteringCampaignTypes = "SEARCH"
)

// Ptr returns reference to report_campaign_get_v2_filtering_campaign_types value
func (v ReportCampaignGetV2FilteringCampaignTypes) Ptr() *ReportCampaignGetV2FilteringCampaignTypes {
	return &v
}
