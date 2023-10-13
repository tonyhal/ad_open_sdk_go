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

// AdlabGroupDetailV30DataDataAdInfoDeepBidType
type AdlabGroupDetailV30DataDataAdInfoDeepBidType string

// List of adlab_group_detail_v3.0_data_data_ad_info_deep_bid_type
const (
	AUTO_MIN_SECOND_STAGE_AdlabGroupDetailV30DataDataAdInfoDeepBidType        AdlabGroupDetailV30DataDataAdInfoDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_AdlabGroupDetailV30DataDataAdInfoDeepBidType               AdlabGroupDetailV30DataDataAdInfoDeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_AdlabGroupDetailV30DataDataAdInfoDeepBidType             AdlabGroupDetailV30DataDataAdInfoDeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_AdlabGroupDetailV30DataDataAdInfoDeepBidType                 AdlabGroupDetailV30DataDataAdInfoDeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_AdlabGroupDetailV30DataDataAdInfoDeepBidType              AdlabGroupDetailV30DataDataAdInfoDeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_AdlabGroupDetailV30DataDataAdInfoDeepBidType AdlabGroupDetailV30DataDataAdInfoDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	MIN_SECOND_STAGE_AdlabGroupDetailV30DataDataAdInfoDeepBidType             AdlabGroupDetailV30DataDataAdInfoDeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_AdlabGroupDetailV30DataDataAdInfoDeepBidType          AdlabGroupDetailV30DataDataAdInfoDeepBidType = "PACING_SECOND_STAGE"
	ROI_COEFFICIENT_AdlabGroupDetailV30DataDataAdInfoDeepBidType              AdlabGroupDetailV30DataDataAdInfoDeepBidType = "ROI_COEFFICIENT"
	ROI_PACING_AdlabGroupDetailV30DataDataAdInfoDeepBidType                   AdlabGroupDetailV30DataDataAdInfoDeepBidType = "ROI_PACING"
	SMARTBID_AdlabGroupDetailV30DataDataAdInfoDeepBidType                     AdlabGroupDetailV30DataDataAdInfoDeepBidType = "SMARTBID"
	SOCIAL_ROI_AdlabGroupDetailV30DataDataAdInfoDeepBidType                   AdlabGroupDetailV30DataDataAdInfoDeepBidType = "SOCIAL_ROI"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoDeepBidType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoDeepBidTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoDeepBidType{
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

// NewAdlabGroupDetailV30DataDataAdInfoDeepBidTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoDeepBidTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoDeepBidType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoDeepBidType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoDeepBidType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_deep_bid_type value
func (v AdlabGroupDetailV30DataDataAdInfoDeepBidType) Ptr() *AdlabGroupDetailV30DataDataAdInfoDeepBidType {
	return &v
}
