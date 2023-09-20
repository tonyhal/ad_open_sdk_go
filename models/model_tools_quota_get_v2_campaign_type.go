/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsQuotaGetV2CampaignType
type ToolsQuotaGetV2CampaignType string

// List of tools_quota_get_v2_campaign_type
const (
	FEED_ToolsQuotaGetV2CampaignType   ToolsQuotaGetV2CampaignType = "FEED"
	SEARCH_ToolsQuotaGetV2CampaignType ToolsQuotaGetV2CampaignType = "SEARCH"
)

// All allowed values of ToolsQuotaGetV2CampaignType enum
var AllowedToolsQuotaGetV2CampaignTypeEnumValues = []ToolsQuotaGetV2CampaignType{
	"FEED",
	"SEARCH",
}

// NewToolsQuotaGetV2CampaignTypeFromValue returns a pointer to a valid ToolsQuotaGetV2CampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsQuotaGetV2CampaignTypeFromValue(v string) (*ToolsQuotaGetV2CampaignType, error) {
	ev := ToolsQuotaGetV2CampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsQuotaGetV2CampaignType: valid values are %v", v, AllowedToolsQuotaGetV2CampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsQuotaGetV2CampaignType) IsValid() bool {
	for _, existing := range AllowedToolsQuotaGetV2CampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_quota_get_v2_campaign_type value
func (v ToolsQuotaGetV2CampaignType) Ptr() *ToolsQuotaGetV2CampaignType {
	return &v
}