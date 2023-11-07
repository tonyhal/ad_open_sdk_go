/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted
type AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_hide_if_converted
const (
	AD_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted           AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "AD"
	ADVERTISER_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted   AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "ADVERTISER"
	APP_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted          AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "APP"
	CAMPAIGN_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted     AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "CAMPAIGN"
	CUSTOMER_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted     AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted   AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted = "ORGANIZATION"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceHideIfConvertedEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted{
	"AD",
	"ADVERTISER",
	"APP",
	"CAMPAIGN",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceHideIfConvertedFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceHideIfConvertedFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_hide_if_converted value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted {
	return &v
}
