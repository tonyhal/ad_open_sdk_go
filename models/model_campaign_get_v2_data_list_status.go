/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2DataListStatus
type CampaignGetV2DataListStatus string

// List of campaign_get_v2_data_list_status
const (
	CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED_CampaignGetV2DataListStatus CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED"
	CAMPAIGN_STATUS_DELETE_CampaignGetV2DataListStatus                   CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_DELETE"
	CAMPAIGN_STATUS_ENABLE_CampaignGetV2DataListStatus                   CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_ENABLE"
	CAMPAIGN_STATUS_DISABLE_CampaignGetV2DataListStatus                  CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_DISABLE"
)

// All allowed values of CampaignGetV2DataListStatus enum
var AllowedCampaignGetV2DataListStatusEnumValues = []CampaignGetV2DataListStatus{
	"CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED",
	"CAMPAIGN_STATUS_DELETE",
	"CAMPAIGN_STATUS_ENABLE",
	"CAMPAIGN_STATUS_DISABLE",
}

// NewCampaignGetV2DataListStatusFromValue returns a pointer to a valid CampaignGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListStatusFromValue(v string) (*CampaignGetV2DataListStatus, error) {
	ev := CampaignGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListStatus: valid values are %v", v, AllowedCampaignGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_status value
func (v CampaignGetV2DataListStatus) Ptr() *CampaignGetV2DataListStatus {
	return &v
}
