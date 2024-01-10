/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignCreateV2LandingType
type CampaignCreateV2LandingType string

// List of campaign_create_v2_landing_type
const (
	GOODS_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "GOODS"
	SHOP_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "SHOP"
	ARTICLE_CampaignCreateV2LandingType        CampaignCreateV2LandingType = "ARTICLE"
	STORE_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "STORE"
	APP_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "APP"
	AWEME_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "AWEME"
	BRAND_EXTERNAL_CampaignCreateV2LandingType CampaignCreateV2LandingType = "BRAND_EXTERNAL"
	LINK_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LINK"
	DPA_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "DPA"
	LIVE_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LIVE"
	QUICK_APP_CampaignCreateV2LandingType      CampaignCreateV2LandingType = "QUICK_APP"
)

// All allowed values of CampaignCreateV2LandingType enum
var AllowedCampaignCreateV2LandingTypeEnumValues = []CampaignCreateV2LandingType{
	"GOODS",
	"SHOP",
	"ARTICLE",
	"STORE",
	"APP",
	"AWEME",
	"BRAND_EXTERNAL",
	"LINK",
	"DPA",
	"LIVE",
	"QUICK_APP",
}

// NewCampaignCreateV2LandingTypeFromValue returns a pointer to a valid CampaignCreateV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2LandingTypeFromValue(v string) (*CampaignCreateV2LandingType, error) {
	ev := CampaignCreateV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2LandingType: valid values are %v", v, AllowedCampaignCreateV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2LandingType) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_landing_type value
func (v CampaignCreateV2LandingType) Ptr() *CampaignCreateV2LandingType {
	return &v
}
