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

// CampaignCreateV2DedicateType
type CampaignCreateV2DedicateType string

// List of campaign_create_v2_dedicate_type
const (
	DEDICATED_CampaignCreateV2DedicateType CampaignCreateV2DedicateType = "DEDICATED"
	UNSET_CampaignCreateV2DedicateType     CampaignCreateV2DedicateType = "UNSET"
)

// All allowed values of CampaignCreateV2DedicateType enum
var AllowedCampaignCreateV2DedicateTypeEnumValues = []CampaignCreateV2DedicateType{
	"DEDICATED",
	"UNSET",
}

// NewCampaignCreateV2DedicateTypeFromValue returns a pointer to a valid CampaignCreateV2DedicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2DedicateTypeFromValue(v string) (*CampaignCreateV2DedicateType, error) {
	ev := CampaignCreateV2DedicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2DedicateType: valid values are %v", v, AllowedCampaignCreateV2DedicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2DedicateType) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2DedicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_dedicate_type value
func (v CampaignCreateV2DedicateType) Ptr() *CampaignCreateV2DedicateType {
	return &v
}
