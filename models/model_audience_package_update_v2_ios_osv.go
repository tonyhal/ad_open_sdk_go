/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2IosOsv
type AudiencePackageUpdateV2IosOsv string

// List of audience_package_update_v2_ios_osv
const (
	Enum_11_2_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "11.2"
	Enum_12_4_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "12.4"
	NONE_AudiencePackageUpdateV2IosOsv      AudiencePackageUpdateV2IosOsv = "NONE"
	Enum_12_1_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "12.1"
	Enum_14_0_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "14.0"
	Enum_4_3_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "4.3"
	Enum_11_0_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "11.0"
	Enum_8_1_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "8.1"
	Enum_8_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "8.0"
	Enum_12_0_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "12.0"
	Enum_12_2_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "12.2"
	Enum_11_1_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "11.1"
	Enum_13_2_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "13.2"
	Enum_4_2_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "4.2"
	Enum_9_3_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "9.3"
	Enum_6_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "6.0"
	Enum_10_3_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "10.3"
	Enum_7_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "7.0"
	Enum_4_1_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "4.1"
	Enum_11_3_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "11.3"
	Enum_10_1_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "10.1"
	Enum_13_4_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "13.4"
	Enum_7_1_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "7.1"
	Enum_12_3_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "12.3"
	Enum_13_3_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "13.3"
	Enum_11_4_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "11.4"
	Enum_4_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "4.0"
	Enum_5_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "5.0"
	Enum_9_2_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "9.2"
	Enum_5_1_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "5.1"
	Enum_8_2_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "8.2"
	Enum_13_0_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "13.0"
	Enum_9_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "9.0"
	Enum_10_2_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "10.2"
	Enum_13_1_AudiencePackageUpdateV2IosOsv AudiencePackageUpdateV2IosOsv = "13.1"
	Enum_0_0_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "0.0"
	Enum_9_1_AudiencePackageUpdateV2IosOsv  AudiencePackageUpdateV2IosOsv = "9.1"
)

// All allowed values of AudiencePackageUpdateV2IosOsv enum
var AllowedAudiencePackageUpdateV2IosOsvEnumValues = []AudiencePackageUpdateV2IosOsv{
	"11.2",
	"12.4",
	"NONE",
	"12.1",
	"14.0",
	"4.3",
	"11.0",
	"8.1",
	"8.0",
	"12.0",
	"12.2",
	"11.1",
	"13.2",
	"4.2",
	"9.3",
	"6.0",
	"10.3",
	"7.0",
	"4.1",
	"11.3",
	"10.1",
	"13.4",
	"7.1",
	"12.3",
	"13.3",
	"11.4",
	"4.0",
	"5.0",
	"9.2",
	"5.1",
	"8.2",
	"13.0",
	"9.0",
	"10.2",
	"13.1",
	"0.0",
	"9.1",
}

// NewAudiencePackageUpdateV2IosOsvFromValue returns a pointer to a valid AudiencePackageUpdateV2IosOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2IosOsvFromValue(v string) (*AudiencePackageUpdateV2IosOsv, error) {
	ev := AudiencePackageUpdateV2IosOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2IosOsv: valid values are %v", v, AllowedAudiencePackageUpdateV2IosOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2IosOsv) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2IosOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_ios_osv value
func (v AudiencePackageUpdateV2IosOsv) Ptr() *AudiencePackageUpdateV2IosOsv {
	return &v
}
