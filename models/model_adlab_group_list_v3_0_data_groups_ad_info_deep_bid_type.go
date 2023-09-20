/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoDeepBidType
type AdlabGroupListV30DataGroupsAdInfoDeepBidType string

// List of adlab_group_list_v3.0_data_groups_ad_info_deep_bid_type
const (
	AUTO_MIN_SECOND_STAGE_AdlabGroupListV30DataGroupsAdInfoDeepBidType        AdlabGroupListV30DataGroupsAdInfoDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_AdlabGroupListV30DataGroupsAdInfoDeepBidType               AdlabGroupListV30DataGroupsAdInfoDeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_AdlabGroupListV30DataGroupsAdInfoDeepBidType             AdlabGroupListV30DataGroupsAdInfoDeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_AdlabGroupListV30DataGroupsAdInfoDeepBidType                 AdlabGroupListV30DataGroupsAdInfoDeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_AdlabGroupListV30DataGroupsAdInfoDeepBidType              AdlabGroupListV30DataGroupsAdInfoDeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_AdlabGroupListV30DataGroupsAdInfoDeepBidType AdlabGroupListV30DataGroupsAdInfoDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	MIN_SECOND_STAGE_AdlabGroupListV30DataGroupsAdInfoDeepBidType             AdlabGroupListV30DataGroupsAdInfoDeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_AdlabGroupListV30DataGroupsAdInfoDeepBidType          AdlabGroupListV30DataGroupsAdInfoDeepBidType = "PACING_SECOND_STAGE"
	ROI_COEFFICIENT_AdlabGroupListV30DataGroupsAdInfoDeepBidType              AdlabGroupListV30DataGroupsAdInfoDeepBidType = "ROI_COEFFICIENT"
	ROI_PACING_AdlabGroupListV30DataGroupsAdInfoDeepBidType                   AdlabGroupListV30DataGroupsAdInfoDeepBidType = "ROI_PACING"
	SMARTBID_AdlabGroupListV30DataGroupsAdInfoDeepBidType                     AdlabGroupListV30DataGroupsAdInfoDeepBidType = "SMARTBID"
	SOCIAL_ROI_AdlabGroupListV30DataGroupsAdInfoDeepBidType                   AdlabGroupListV30DataGroupsAdInfoDeepBidType = "SOCIAL_ROI"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoDeepBidType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoDeepBidTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoDeepBidType{
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"ROI_COEFFICIENT",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
}

// NewAdlabGroupListV30DataGroupsAdInfoDeepBidTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoDeepBidTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoDeepBidType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoDeepBidType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoDeepBidType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_deep_bid_type value
func (v AdlabGroupListV30DataGroupsAdInfoDeepBidType) Ptr() *AdlabGroupListV30DataGroupsAdInfoDeepBidType {
	return &v
}
