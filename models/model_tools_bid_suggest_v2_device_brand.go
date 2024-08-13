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

// ToolsBidSuggestV2DeviceBrand
type ToolsBidSuggestV2DeviceBrand string

// List of tools_bid_suggest_v2_device_brand
const (
	APPLE_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "APPLE"
	GOOGLE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GOOGLE"
	PEPPER_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "PEPPER"
	NOKIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NOKIA"
	SAMSUNG_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "SAMSUNG"
	ZTE_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "ZTE"
	COOLPAD_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "COOLPAD"
	HISENSE_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "HISENSE"
	HUAWEI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "HUAWEI"
	CHINAMOBILE_ToolsBidSuggestV2DeviceBrand ToolsBidSuggestV2DeviceBrand = "CHINAMOBILE"
	ONEPLUS_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "ONEPLUS"
	NUBIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NUBIA"
	LG_ToolsBidSuggestV2DeviceBrand          ToolsBidSuggestV2DeviceBrand = "LG"
	SONY_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "SONY"
	Enum_360_ToolsBidSuggestV2DeviceBrand    ToolsBidSuggestV2DeviceBrand = "360"
	HONOR_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "HONOR"
	LETV_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "LETV"
	OPPO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "OPPO"
	SMARTISAN_ToolsBidSuggestV2DeviceBrand   ToolsBidSuggestV2DeviceBrand = "SMARTISAN"
	VIVO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "VIVO"
	TCL_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "TCL"
	LENOVO_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "LENOVO"
	XIAOMI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "XIAOMI"
	MEIZU_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "MEIZU"
	QIKU_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "QIKU"
	MOTO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "MOTO"
	HTC_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "HTC"
	GIONEE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GIONEE"
)

// All allowed values of ToolsBidSuggestV2DeviceBrand enum
var AllowedToolsBidSuggestV2DeviceBrandEnumValues = []ToolsBidSuggestV2DeviceBrand{
	"APPLE",
	"GOOGLE",
	"PEPPER",
	"NOKIA",
	"SAMSUNG",
	"ZTE",
	"COOLPAD",
	"HISENSE",
	"HUAWEI",
	"CHINAMOBILE",
	"ONEPLUS",
	"NUBIA",
	"LG",
	"SONY",
	"360",
	"HONOR",
	"LETV",
	"OPPO",
	"SMARTISAN",
	"VIVO",
	"TCL",
	"LENOVO",
	"XIAOMI",
	"MEIZU",
	"QIKU",
	"MOTO",
	"HTC",
	"GIONEE",
}

// NewToolsBidSuggestV2DeviceBrandFromValue returns a pointer to a valid ToolsBidSuggestV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DeviceBrandFromValue(v string) (*ToolsBidSuggestV2DeviceBrand, error) {
	ev := ToolsBidSuggestV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2DeviceBrand: valid values are %v", v, AllowedToolsBidSuggestV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_device_brand value
func (v ToolsBidSuggestV2DeviceBrand) Ptr() *ToolsBidSuggestV2DeviceBrand {
	return &v
}
