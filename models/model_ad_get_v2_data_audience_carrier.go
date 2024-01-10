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

// AdGetV2DataAudienceCarrier
type AdGetV2DataAudienceCarrier string

// List of ad_get_v2_data_audience_carrier
const (
	MOBILE_AdGetV2DataAudienceCarrier AdGetV2DataAudienceCarrier = "MOBILE"
	TELCOM_AdGetV2DataAudienceCarrier AdGetV2DataAudienceCarrier = "TELCOM"
	UNICOM_AdGetV2DataAudienceCarrier AdGetV2DataAudienceCarrier = "UNICOM"
)

// All allowed values of AdGetV2DataAudienceCarrier enum
var AllowedAdGetV2DataAudienceCarrierEnumValues = []AdGetV2DataAudienceCarrier{
	"MOBILE",
	"TELCOM",
	"UNICOM",
}

// NewAdGetV2DataAudienceCarrierFromValue returns a pointer to a valid AdGetV2DataAudienceCarrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceCarrierFromValue(v string) (*AdGetV2DataAudienceCarrier, error) {
	ev := AdGetV2DataAudienceCarrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceCarrier: valid values are %v", v, AllowedAdGetV2DataAudienceCarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceCarrier) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceCarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_carrier value
func (v AdGetV2DataAudienceCarrier) Ptr() *AdGetV2DataAudienceCarrier {
	return &v
}
