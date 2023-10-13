/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2HideIfConverted
type AudiencePackageCreateV2HideIfConverted string

// List of audience_package_create_v2_hide_if_converted
const (
	AD_AudiencePackageCreateV2HideIfConverted           AudiencePackageCreateV2HideIfConverted = "AD"
	CAMPAIGN_AudiencePackageCreateV2HideIfConverted     AudiencePackageCreateV2HideIfConverted = "CAMPAIGN"
	ADVERTISER_AudiencePackageCreateV2HideIfConverted   AudiencePackageCreateV2HideIfConverted = "ADVERTISER"
	ORGANIZATION_AudiencePackageCreateV2HideIfConverted AudiencePackageCreateV2HideIfConverted = "ORGANIZATION"
	NO_EXCLUDE_AudiencePackageCreateV2HideIfConverted   AudiencePackageCreateV2HideIfConverted = "NO_EXCLUDE"
	CUSTOMER_AudiencePackageCreateV2HideIfConverted     AudiencePackageCreateV2HideIfConverted = "CUSTOMER"
	APP_AudiencePackageCreateV2HideIfConverted          AudiencePackageCreateV2HideIfConverted = "APP"
)

// All allowed values of AudiencePackageCreateV2HideIfConverted enum
var AllowedAudiencePackageCreateV2HideIfConvertedEnumValues = []AudiencePackageCreateV2HideIfConverted{
	"AD",
	"CAMPAIGN",
	"ADVERTISER",
	"ORGANIZATION",
	"NO_EXCLUDE",
	"CUSTOMER",
	"APP",
}

// NewAudiencePackageCreateV2HideIfConvertedFromValue returns a pointer to a valid AudiencePackageCreateV2HideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2HideIfConvertedFromValue(v string) (*AudiencePackageCreateV2HideIfConverted, error) {
	ev := AudiencePackageCreateV2HideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2HideIfConverted: valid values are %v", v, AllowedAudiencePackageCreateV2HideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2HideIfConverted) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2HideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_hide_if_converted value
func (v AudiencePackageCreateV2HideIfConverted) Ptr() *AudiencePackageCreateV2HideIfConverted {
	return &v
}
