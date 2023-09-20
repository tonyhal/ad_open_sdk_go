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

// AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum
type AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum string

// List of adlab_group_list_v3.0_data_groups_campaign_info_delivery_related_num
const (
	CAMPAIGN_DPA_DEFAULT_NOT_AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum     AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum = "CAMPAIGN_DPA_DEFAULT_NOT"
	CAMPAIGN_DPA_SINGLE_DELIVERY_AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum = "CAMPAIGN_DPA_SINGLE_DELIVERY"
)

// All allowed values of AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum enum
var AllowedAdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNumEnumValues = []AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum{
	"CAMPAIGN_DPA_DEFAULT_NOT",
	"CAMPAIGN_DPA_SINGLE_DELIVERY",
}

// NewAdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNumFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNumFromValue(v string) (*AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum, error) {
	ev := AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_campaign_info_delivery_related_num value
func (v AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum) Ptr() *AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum {
	return &v
}
