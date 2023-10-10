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

// ToolsBidSuggestV2DeviceBrand
type ToolsBidSuggestV2DeviceBrand string

// List of tools_bid_suggest_v2_device_brand
const (
	NUBIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NUBIA"
	XIAOMI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "XIAOMI"
	VIVO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "VIVO"
	GOOGLE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GOOGLE"
	LETV_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "LETV"
	HTC_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "HTC"
	ZTE_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "ZTE"
	ONEPLUS_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "ONEPLUS"
	SMARTISAN_ToolsBidSuggestV2DeviceBrand   ToolsBidSuggestV2DeviceBrand = "SMARTISAN"
	Enum_360_ToolsBidSuggestV2DeviceBrand    ToolsBidSuggestV2DeviceBrand = "360"
	MOTO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "MOTO"
	SAMSUNG_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "SAMSUNG"
	LG_ToolsBidSuggestV2DeviceBrand          ToolsBidSuggestV2DeviceBrand = "LG"
	TCL_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "TCL"
	APPLE_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "APPLE"
	HONOR_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "HONOR"
	NOKIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NOKIA"
	CHINAMOBILE_ToolsBidSuggestV2DeviceBrand ToolsBidSuggestV2DeviceBrand = "CHINAMOBILE"
	MEIZU_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "MEIZU"
	SONY_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "SONY"
	PEPPER_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "PEPPER"
	HISENSE_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "HISENSE"
	OPPO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "OPPO"
	HUAWEI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "HUAWEI"
	COOLPAD_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "COOLPAD"
	LENOVO_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "LENOVO"
	GIONEE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GIONEE"
	QIKU_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "QIKU"
)

// All allowed values of ToolsBidSuggestV2DeviceBrand enum
var AllowedToolsBidSuggestV2DeviceBrandEnumValues = []ToolsBidSuggestV2DeviceBrand{
	"NUBIA",
	"XIAOMI",
	"VIVO",
	"GOOGLE",
	"LETV",
	"HTC",
	"ZTE",
	"ONEPLUS",
	"SMARTISAN",
	"360",
	"MOTO",
	"SAMSUNG",
	"LG",
	"TCL",
	"APPLE",
	"HONOR",
	"NOKIA",
	"CHINAMOBILE",
	"MEIZU",
	"SONY",
	"PEPPER",
	"HISENSE",
	"OPPO",
	"HUAWEI",
	"COOLPAD",
	"LENOVO",
	"GIONEE",
	"QIKU",
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
