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

// AdGetV2DataAudienceDpaRtaSwitch
type AdGetV2DataAudienceDpaRtaSwitch string

// List of ad_get_v2_data_audience_dpa_rta_switch
const (
	OFF_AdGetV2DataAudienceDpaRtaSwitch AdGetV2DataAudienceDpaRtaSwitch = "OFF"
	ON_AdGetV2DataAudienceDpaRtaSwitch  AdGetV2DataAudienceDpaRtaSwitch = "ON"
)

// All allowed values of AdGetV2DataAudienceDpaRtaSwitch enum
var AllowedAdGetV2DataAudienceDpaRtaSwitchEnumValues = []AdGetV2DataAudienceDpaRtaSwitch{
	"OFF",
	"ON",
}

// NewAdGetV2DataAudienceDpaRtaSwitchFromValue returns a pointer to a valid AdGetV2DataAudienceDpaRtaSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDpaRtaSwitchFromValue(v string) (*AdGetV2DataAudienceDpaRtaSwitch, error) {
	ev := AdGetV2DataAudienceDpaRtaSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDpaRtaSwitch: valid values are %v", v, AllowedAdGetV2DataAudienceDpaRtaSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDpaRtaSwitch) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDpaRtaSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_dpa_rta_switch value
func (v AdGetV2DataAudienceDpaRtaSwitch) Ptr() *AdGetV2DataAudienceDpaRtaSwitch {
	return &v
}
