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

// CampaignGetV2DataListDeliveryRelatedNum
type CampaignGetV2DataListDeliveryRelatedNum string

// List of campaign_get_v2_data_list_delivery_related_num
const (
	CAMPAIGN_DPA_SINGLE_DELIVERY_CampaignGetV2DataListDeliveryRelatedNum CampaignGetV2DataListDeliveryRelatedNum = "CAMPAIGN_DPA_SINGLE_DELIVERY"
	CAMPAIGN_DPA_MULTI_DELIVERY_CampaignGetV2DataListDeliveryRelatedNum  CampaignGetV2DataListDeliveryRelatedNum = "CAMPAIGN_DPA_MULTI_DELIVERY"
	CAMPAIGN_DPA_DEFAULT_NOT_CampaignGetV2DataListDeliveryRelatedNum     CampaignGetV2DataListDeliveryRelatedNum = "CAMPAIGN_DPA_DEFAULT_NOT"
)

// All allowed values of CampaignGetV2DataListDeliveryRelatedNum enum
var AllowedCampaignGetV2DataListDeliveryRelatedNumEnumValues = []CampaignGetV2DataListDeliveryRelatedNum{
	"CAMPAIGN_DPA_SINGLE_DELIVERY",
	"CAMPAIGN_DPA_MULTI_DELIVERY",
	"CAMPAIGN_DPA_DEFAULT_NOT",
}

// NewCampaignGetV2DataListDeliveryRelatedNumFromValue returns a pointer to a valid CampaignGetV2DataListDeliveryRelatedNum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListDeliveryRelatedNumFromValue(v string) (*CampaignGetV2DataListDeliveryRelatedNum, error) {
	ev := CampaignGetV2DataListDeliveryRelatedNum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListDeliveryRelatedNum: valid values are %v", v, AllowedCampaignGetV2DataListDeliveryRelatedNumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListDeliveryRelatedNum) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListDeliveryRelatedNumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_delivery_related_num value
func (v CampaignGetV2DataListDeliveryRelatedNum) Ptr() *CampaignGetV2DataListDeliveryRelatedNum {
	return &v
}
