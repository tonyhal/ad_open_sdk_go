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

// AdGetV2DataAudienceAndroidOsv
type AdGetV2DataAudienceAndroidOsv string

// List of ad_get_v2_data_audience_android_osv
const (
	Enum_2_1_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "2.1"
	Enum_2_2_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "2.2"
	Enum_11_0_AdGetV2DataAudienceAndroidOsv AdGetV2DataAudienceAndroidOsv = "11.0"
	NONE_AdGetV2DataAudienceAndroidOsv      AdGetV2DataAudienceAndroidOsv = "NONE"
	Enum_2_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "2.0"
	Enum_4_2_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "4.2"
	Enum_4_3_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "4.3"
	Enum_2_3_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "2.3"
	Enum_3_1_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "3.1"
	Enum_10_1_AdGetV2DataAudienceAndroidOsv AdGetV2DataAudienceAndroidOsv = "10.1"
	Enum_4_4_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "4.4"
	Enum_3_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "3.0"
	Enum_7_1_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "7.1"
	Enum_8_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "8.0"
	Enum_0_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "0.0"
	Enum_4_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "4.0"
	Enum_5_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "5.0"
	Enum_9_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "9.0"
	Enum_6_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "6.0"
	Enum_8_1_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "8.1"
	Enum_5_1_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "5.1"
	Enum_4_5_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "4.5"
	Enum_7_0_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "7.0"
	Enum_10_0_AdGetV2DataAudienceAndroidOsv AdGetV2DataAudienceAndroidOsv = "10.0"
	Enum_3_2_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "3.2"
	Enum_4_1_AdGetV2DataAudienceAndroidOsv  AdGetV2DataAudienceAndroidOsv = "4.1"
)

// All allowed values of AdGetV2DataAudienceAndroidOsv enum
var AllowedAdGetV2DataAudienceAndroidOsvEnumValues = []AdGetV2DataAudienceAndroidOsv{
	"2.1",
	"2.2",
	"11.0",
	"NONE",
	"2.0",
	"4.2",
	"4.3",
	"2.3",
	"3.1",
	"10.1",
	"4.4",
	"3.0",
	"7.1",
	"8.0",
	"0.0",
	"4.0",
	"5.0",
	"9.0",
	"6.0",
	"8.1",
	"5.1",
	"4.5",
	"7.0",
	"10.0",
	"3.2",
	"4.1",
}

// NewAdGetV2DataAudienceAndroidOsvFromValue returns a pointer to a valid AdGetV2DataAudienceAndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAndroidOsvFromValue(v string) (*AdGetV2DataAudienceAndroidOsv, error) {
	ev := AdGetV2DataAudienceAndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAndroidOsv: valid values are %v", v, AllowedAdGetV2DataAudienceAndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAndroidOsv) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_android_osv value
func (v AdGetV2DataAudienceAndroidOsv) Ptr() *AdGetV2DataAudienceAndroidOsv {
	return &v
}
