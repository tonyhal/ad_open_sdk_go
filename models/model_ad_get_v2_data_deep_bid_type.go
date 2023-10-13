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

// AdGetV2DataDeepBidType
type AdGetV2DataDeepBidType string

// List of ad_get_v2_data_deep_bid_type
const (
	PACING_SECOND_STAGE_AdGetV2DataDeepBidType          AdGetV2DataDeepBidType = "PACING_SECOND_STAGE"
	AUTO_MIN_SECOND_STAGE_AdGetV2DataDeepBidType        AdGetV2DataDeepBidType = "AUTO_MIN_SECOND_STAGE"
	ROI_PACING_AdGetV2DataDeepBidType                   AdGetV2DataDeepBidType = "ROI_PACING"
	BID_PER_ACTION_AdGetV2DataDeepBidType               AdGetV2DataDeepBidType = "BID_PER_ACTION"
	DEEP_BID_TYPE_RETENTION_DAYS_AdGetV2DataDeepBidType AdGetV2DataDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	ROI_COEFFICIENT_AdGetV2DataDeepBidType              AdGetV2DataDeepBidType = "ROI_COEFFICIENT"
	DEEP_BID_MIN_AdGetV2DataDeepBidType                 AdGetV2DataDeepBidType = "DEEP_BID_MIN"
	SMARTBID_AdGetV2DataDeepBidType                     AdGetV2DataDeepBidType = "SMARTBID"
	ROI_DIRECT_MAIL_AdGetV2DataDeepBidType              AdGetV2DataDeepBidType = "ROI_DIRECT_MAIL"
	DEEP_BID_DEFAULT_AdGetV2DataDeepBidType             AdGetV2DataDeepBidType = "DEEP_BID_DEFAULT"
	SOCIAL_ROI_AdGetV2DataDeepBidType                   AdGetV2DataDeepBidType = "SOCIAL_ROI"
	MIN_SECOND_STAGE_AdGetV2DataDeepBidType             AdGetV2DataDeepBidType = "MIN_SECOND_STAGE"
	ALI_FL_AdGetV2DataDeepBidType                       AdGetV2DataDeepBidType = "ALI_FL"
	DEEP_BID_PACING_AdGetV2DataDeepBidType              AdGetV2DataDeepBidType = "DEEP_BID_PACING"
)

// All allowed values of AdGetV2DataDeepBidType enum
var AllowedAdGetV2DataDeepBidTypeEnumValues = []AdGetV2DataDeepBidType{
	"PACING_SECOND_STAGE",
	"AUTO_MIN_SECOND_STAGE",
	"ROI_PACING",
	"BID_PER_ACTION",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"ROI_COEFFICIENT",
	"DEEP_BID_MIN",
	"SMARTBID",
	"ROI_DIRECT_MAIL",
	"DEEP_BID_DEFAULT",
	"SOCIAL_ROI",
	"MIN_SECOND_STAGE",
	"ALI_FL",
	"DEEP_BID_PACING",
}

// NewAdGetV2DataDeepBidTypeFromValue returns a pointer to a valid AdGetV2DataDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDeepBidTypeFromValue(v string) (*AdGetV2DataDeepBidType, error) {
	ev := AdGetV2DataDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDeepBidType: valid values are %v", v, AllowedAdGetV2DataDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDeepBidType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_deep_bid_type value
func (v AdGetV2DataDeepBidType) Ptr() *AdGetV2DataDeepBidType {
	return &v
}
