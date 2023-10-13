/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2DataListDeliveryMode
type CampaignGetV2DataListDeliveryMode string

// List of campaign_get_v2_data_list_delivery_mode
const (
	MANUAL_CampaignGetV2DataListDeliveryMode     CampaignGetV2DataListDeliveryMode = "MANUAL"
	PROCEDURAL_CampaignGetV2DataListDeliveryMode CampaignGetV2DataListDeliveryMode = "PROCEDURAL"
)

// All allowed values of CampaignGetV2DataListDeliveryMode enum
var AllowedCampaignGetV2DataListDeliveryModeEnumValues = []CampaignGetV2DataListDeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewCampaignGetV2DataListDeliveryModeFromValue returns a pointer to a valid CampaignGetV2DataListDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListDeliveryModeFromValue(v string) (*CampaignGetV2DataListDeliveryMode, error) {
	ev := CampaignGetV2DataListDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListDeliveryMode: valid values are %v", v, AllowedCampaignGetV2DataListDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListDeliveryMode) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_delivery_mode value
func (v CampaignGetV2DataListDeliveryMode) Ptr() *CampaignGetV2DataListDeliveryMode {
	return &v
}
