/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2IosOsv
type ToolsEstimateAudienceV2IosOsv string

// List of tools_estimate_audience_v2_ios_osv
const (
	Enum_8_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "8.0"
	Enum_12_4_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "12.4"
	Enum_5_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "5.0"
	Enum_7_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "7.0"
	Enum_4_3_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "4.3"
	Enum_9_1_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "9.1"
	Enum_11_2_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "11.2"
	Enum_14_0_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "14.0"
	Enum_7_1_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "7.1"
	Enum_5_1_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "5.1"
	Enum_12_1_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "12.1"
	Enum_6_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "6.0"
	Enum_13_1_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "13.1"
	Enum_8_2_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "8.2"
	Enum_8_1_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "8.1"
	Enum_13_2_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "13.2"
	Enum_12_2_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "12.2"
	NONE_ToolsEstimateAudienceV2IosOsv      ToolsEstimateAudienceV2IosOsv = "NONE"
	Enum_4_2_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "4.2"
	Enum_4_1_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "4.1"
	Enum_13_4_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "13.4"
	Enum_12_0_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "12.0"
	Enum_11_4_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "11.4"
	Enum_12_3_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "12.3"
	Enum_13_3_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "13.3"
	Enum_0_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "0.0"
	Enum_13_0_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "13.0"
	Enum_9_3_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "9.3"
	Enum_4_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "4.0"
	Enum_11_3_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "11.3"
	Enum_10_1_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "10.1"
	Enum_11_0_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "11.0"
	Enum_9_0_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "9.0"
	Enum_10_2_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "10.2"
	Enum_10_3_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "10.3"
	Enum_11_1_ToolsEstimateAudienceV2IosOsv ToolsEstimateAudienceV2IosOsv = "11.1"
	Enum_9_2_ToolsEstimateAudienceV2IosOsv  ToolsEstimateAudienceV2IosOsv = "9.2"
)

// All allowed values of ToolsEstimateAudienceV2IosOsv enum
var AllowedToolsEstimateAudienceV2IosOsvEnumValues = []ToolsEstimateAudienceV2IosOsv{
	"8.0",
	"12.4",
	"5.0",
	"7.0",
	"4.3",
	"9.1",
	"11.2",
	"14.0",
	"7.1",
	"5.1",
	"12.1",
	"6.0",
	"13.1",
	"8.2",
	"8.1",
	"13.2",
	"12.2",
	"NONE",
	"4.2",
	"4.1",
	"13.4",
	"12.0",
	"11.4",
	"12.3",
	"13.3",
	"0.0",
	"13.0",
	"9.3",
	"4.0",
	"11.3",
	"10.1",
	"11.0",
	"9.0",
	"10.2",
	"10.3",
	"11.1",
	"9.2",
}

// NewToolsEstimateAudienceV2IosOsvFromValue returns a pointer to a valid ToolsEstimateAudienceV2IosOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2IosOsvFromValue(v string) (*ToolsEstimateAudienceV2IosOsv, error) {
	ev := ToolsEstimateAudienceV2IosOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2IosOsv: valid values are %v", v, AllowedToolsEstimateAudienceV2IosOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2IosOsv) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2IosOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_ios_osv value
func (v ToolsEstimateAudienceV2IosOsv) Ptr() *ToolsEstimateAudienceV2IosOsv {
	return &v
}
