/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceIosOsv
type AdGetV2DataAudienceIosOsv string

// List of ad_get_v2_data_audience_ios_osv
const (
	Enum_8_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "8.0"
	Enum_12_4_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "12.4"
	Enum_5_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "5.0"
	Enum_7_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "7.0"
	Enum_4_3_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "4.3"
	Enum_9_1_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "9.1"
	Enum_11_2_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "11.2"
	Enum_14_0_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "14.0"
	Enum_7_1_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "7.1"
	Enum_5_1_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "5.1"
	Enum_12_1_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "12.1"
	Enum_6_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "6.0"
	Enum_13_1_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "13.1"
	Enum_8_2_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "8.2"
	Enum_8_1_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "8.1"
	Enum_13_2_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "13.2"
	Enum_12_2_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "12.2"
	NONE_AdGetV2DataAudienceIosOsv      AdGetV2DataAudienceIosOsv = "NONE"
	Enum_4_2_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "4.2"
	Enum_4_1_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "4.1"
	Enum_13_4_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "13.4"
	Enum_12_0_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "12.0"
	Enum_11_4_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "11.4"
	Enum_12_3_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "12.3"
	Enum_13_3_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "13.3"
	Enum_0_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "0.0"
	Enum_13_0_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "13.0"
	Enum_9_3_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "9.3"
	Enum_4_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "4.0"
	Enum_11_3_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "11.3"
	Enum_10_1_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "10.1"
	Enum_11_0_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "11.0"
	Enum_9_0_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "9.0"
	Enum_10_2_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "10.2"
	Enum_10_3_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "10.3"
	Enum_11_1_AdGetV2DataAudienceIosOsv AdGetV2DataAudienceIosOsv = "11.1"
	Enum_9_2_AdGetV2DataAudienceIosOsv  AdGetV2DataAudienceIosOsv = "9.2"
)

// All allowed values of AdGetV2DataAudienceIosOsv enum
var AllowedAdGetV2DataAudienceIosOsvEnumValues = []AdGetV2DataAudienceIosOsv{
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

// NewAdGetV2DataAudienceIosOsvFromValue returns a pointer to a valid AdGetV2DataAudienceIosOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceIosOsvFromValue(v string) (*AdGetV2DataAudienceIosOsv, error) {
	ev := AdGetV2DataAudienceIosOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceIosOsv: valid values are %v", v, AllowedAdGetV2DataAudienceIosOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceIosOsv) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceIosOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_ios_osv value
func (v AdGetV2DataAudienceIosOsv) Ptr() *AdGetV2DataAudienceIosOsv {
	return &v
}
