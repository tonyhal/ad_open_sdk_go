/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2FilteringDeliveryMode
type ReportCampaignGetV2FilteringDeliveryMode string

// List of report_campaign_get_v2_filtering_delivery_mode
const (
	ADLAB_FREE_ReportCampaignGetV2FilteringDeliveryMode ReportCampaignGetV2FilteringDeliveryMode = "ADLAB_FREE"
	STANDARD_ReportCampaignGetV2FilteringDeliveryMode   ReportCampaignGetV2FilteringDeliveryMode = "STANDARD"
)

// Ptr returns reference to report_campaign_get_v2_filtering_delivery_mode value
func (v ReportCampaignGetV2FilteringDeliveryMode) Ptr() *ReportCampaignGetV2FilteringDeliveryMode {
	return &v
}
