/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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
	AWEME_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "AWEME"
	APP_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "APP"
	SHOP_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "SHOP"
	QUICK_APP_CampaignCreateV2LandingType      CampaignCreateV2LandingType = "QUICK_APP"
	LINK_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LINK"
	DPA_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "DPA"
	ARTICLE_CampaignCreateV2LandingType        CampaignCreateV2LandingType = "ARTICLE"
	GOODS_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "GOODS"
	STORE_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "STORE"
	LIVE_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LIVE"
	BRAND_EXTERNAL_CampaignCreateV2LandingType CampaignCreateV2LandingType = "BRAND_EXTERNAL"
)

// All allowed values of CampaignCreateV2LandingType enum
var AllowedCampaignCreateV2LandingTypeEnumValues = []CampaignCreateV2LandingType{
	"AWEME",
	"APP",
	"SHOP",
	"QUICK_APP",
	"LINK",
	"DPA",
	"ARTICLE",
	"GOODS",
	"STORE",
	"LIVE",
	"BRAND_EXTERNAL",
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