/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandCampaignGetV30DataCampaignsCampaignStatus
type BrandCampaignGetV30DataCampaignsCampaignStatus int64

// List of brand_campaign_get_v3.0_data_campaigns_campaign_status
const (
	Enum_1_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 1
	Enum_10_BrandCampaignGetV30DataCampaignsCampaignStatus BrandCampaignGetV30DataCampaignsCampaignStatus = 10
	Enum_11_BrandCampaignGetV30DataCampaignsCampaignStatus BrandCampaignGetV30DataCampaignsCampaignStatus = 11
	Enum_2_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 2
	Enum_3_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 3
	Enum_4_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 4
	Enum_5_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 5
	Enum_6_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 6
	Enum_7_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 7
	Enum_8_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 8
	Enum_9_BrandCampaignGetV30DataCampaignsCampaignStatus  BrandCampaignGetV30DataCampaignsCampaignStatus = 9
)

// All allowed values of BrandCampaignGetV30DataCampaignsCampaignStatus enum
var AllowedBrandCampaignGetV30DataCampaignsCampaignStatusEnumValues = []BrandCampaignGetV30DataCampaignsCampaignStatus{
	1,
	10,
	11,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewBrandCampaignGetV30DataCampaignsCampaignStatusFromValue returns a pointer to a valid BrandCampaignGetV30DataCampaignsCampaignStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCampaignGetV30DataCampaignsCampaignStatusFromValue(v int64) (*BrandCampaignGetV30DataCampaignsCampaignStatus, error) {
	ev := BrandCampaignGetV30DataCampaignsCampaignStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCampaignGetV30DataCampaignsCampaignStatus: valid values are %v", v, AllowedBrandCampaignGetV30DataCampaignsCampaignStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCampaignGetV30DataCampaignsCampaignStatus) IsValid() bool {
	for _, existing := range AllowedBrandCampaignGetV30DataCampaignsCampaignStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_campaign_get_v3.0_data_campaigns_campaign_status value
func (v BrandCampaignGetV30DataCampaignsCampaignStatus) Ptr() *BrandCampaignGetV30DataCampaignsCampaignStatus {
	return &v
}
