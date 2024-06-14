/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2FilteringDeliveryMode
type ReportCampaignGetV2FilteringDeliveryMode string

// List of report_campaign_get_v2_filtering_delivery_mode
const (
	ADLAB_FREE_ReportCampaignGetV2FilteringDeliveryMode ReportCampaignGetV2FilteringDeliveryMode = "ADLAB_FREE"
	STANDARD_ReportCampaignGetV2FilteringDeliveryMode   ReportCampaignGetV2FilteringDeliveryMode = "STANDARD"
)

// All allowed values of ReportCampaignGetV2FilteringDeliveryMode enum
var AllowedReportCampaignGetV2FilteringDeliveryModeEnumValues = []ReportCampaignGetV2FilteringDeliveryMode{
	"ADLAB_FREE",
	"STANDARD",
}

// NewReportCampaignGetV2FilteringDeliveryModeFromValue returns a pointer to a valid ReportCampaignGetV2FilteringDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringDeliveryModeFromValue(v string) (*ReportCampaignGetV2FilteringDeliveryMode, error) {
	ev := ReportCampaignGetV2FilteringDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringDeliveryMode: valid values are %v", v, AllowedReportCampaignGetV2FilteringDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringDeliveryMode) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_delivery_mode value
func (v ReportCampaignGetV2FilteringDeliveryMode) Ptr() *ReportCampaignGetV2FilteringDeliveryMode {
	return &v
}
