/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceDeviceBrand
type AdGetV2DataAudienceDeviceBrand string

// List of ad_get_v2_data_audience_device_brand
const (
	NUBIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NUBIA"
	QIKU_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "QIKU"
	MOTO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "MOTO"
	SAMSUNG_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "SAMSUNG"
	OPPO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "OPPO"
	GOOGLE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GOOGLE"
	SMARTISAN_AdGetV2DataAudienceDeviceBrand   AdGetV2DataAudienceDeviceBrand = "SMARTISAN"
	LENOVO_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "LENOVO"
	VIVO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "VIVO"
	Enum_360_AdGetV2DataAudienceDeviceBrand    AdGetV2DataAudienceDeviceBrand = "360"
	XIAOMI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "XIAOMI"
	LETV_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "LETV"
	PEPPER_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "PEPPER"
	HTC_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "HTC"
	GIONEE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GIONEE"
	MEIZU_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "MEIZU"
	TCL_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "TCL"
	ONEPLUS_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "ONEPLUS"
	HONOR_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "HONOR"
	HISENSE_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "HISENSE"
	HUAWEI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "HUAWEI"
	APPLE_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "APPLE"
	ZTE_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "ZTE"
	LG_AdGetV2DataAudienceDeviceBrand          AdGetV2DataAudienceDeviceBrand = "LG"
	COOLPAD_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "COOLPAD"
	CHINAMOBILE_AdGetV2DataAudienceDeviceBrand AdGetV2DataAudienceDeviceBrand = "CHINAMOBILE"
	SONY_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "SONY"
	NOKIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NOKIA"
)

// All allowed values of AdGetV2DataAudienceDeviceBrand enum
var AllowedAdGetV2DataAudienceDeviceBrandEnumValues = []AdGetV2DataAudienceDeviceBrand{
	"NUBIA",
	"QIKU",
	"MOTO",
	"SAMSUNG",
	"OPPO",
	"GOOGLE",
	"SMARTISAN",
	"LENOVO",
	"VIVO",
	"360",
	"XIAOMI",
	"LETV",
	"PEPPER",
	"HTC",
	"GIONEE",
	"MEIZU",
	"TCL",
	"ONEPLUS",
	"HONOR",
	"HISENSE",
	"HUAWEI",
	"APPLE",
	"ZTE",
	"LG",
	"COOLPAD",
	"CHINAMOBILE",
	"SONY",
	"NOKIA",
}

// NewAdGetV2DataAudienceDeviceBrandFromValue returns a pointer to a valid AdGetV2DataAudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDeviceBrandFromValue(v string) (*AdGetV2DataAudienceDeviceBrand, error) {
	ev := AdGetV2DataAudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDeviceBrand: valid values are %v", v, AllowedAdGetV2DataAudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_device_brand value
func (v AdGetV2DataAudienceDeviceBrand) Ptr() *AdGetV2DataAudienceDeviceBrand {
	return &v
}
