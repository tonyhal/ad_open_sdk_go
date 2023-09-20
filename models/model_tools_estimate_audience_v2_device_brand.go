/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2DeviceBrand
type ToolsEstimateAudienceV2DeviceBrand string

// List of tools_estimate_audience_v2_device_brand
const (
	MEIZU_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "MEIZU"
	GIONEE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GIONEE"
	CHINAMOBILE_ToolsEstimateAudienceV2DeviceBrand ToolsEstimateAudienceV2DeviceBrand = "CHINAMOBILE"
	OPPO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "OPPO"
	HUAWEI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "HUAWEI"
	GOOGLE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GOOGLE"
	HISENSE_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "HISENSE"
	ZTE_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "ZTE"
	SONY_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "SONY"
	COOLPAD_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "COOLPAD"
	VIVO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "VIVO"
	XIAOMI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "XIAOMI"
	SMARTISAN_ToolsEstimateAudienceV2DeviceBrand   ToolsEstimateAudienceV2DeviceBrand = "SMARTISAN"
	NUBIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NUBIA"
	LENOVO_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "LENOVO"
	TCL_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "TCL"
	LG_ToolsEstimateAudienceV2DeviceBrand          ToolsEstimateAudienceV2DeviceBrand = "LG"
	NOKIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NOKIA"
	PEPPER_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "PEPPER"
	QIKU_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "QIKU"
	HONOR_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "HONOR"
	Enum_360_ToolsEstimateAudienceV2DeviceBrand    ToolsEstimateAudienceV2DeviceBrand = "360"
	HTC_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "HTC"
	SAMSUNG_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "SAMSUNG"
	LETV_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "LETV"
	MOTO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "MOTO"
	ONEPLUS_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "ONEPLUS"
	APPLE_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "APPLE"
)

// All allowed values of ToolsEstimateAudienceV2DeviceBrand enum
var AllowedToolsEstimateAudienceV2DeviceBrandEnumValues = []ToolsEstimateAudienceV2DeviceBrand{
	"MEIZU",
	"GIONEE",
	"CHINAMOBILE",
	"OPPO",
	"HUAWEI",
	"GOOGLE",
	"HISENSE",
	"ZTE",
	"SONY",
	"COOLPAD",
	"VIVO",
	"XIAOMI",
	"SMARTISAN",
	"NUBIA",
	"LENOVO",
	"TCL",
	"LG",
	"NOKIA",
	"PEPPER",
	"QIKU",
	"HONOR",
	"360",
	"HTC",
	"SAMSUNG",
	"LETV",
	"MOTO",
	"ONEPLUS",
	"APPLE",
}

// NewToolsEstimateAudienceV2DeviceBrandFromValue returns a pointer to a valid ToolsEstimateAudienceV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2DeviceBrandFromValue(v string) (*ToolsEstimateAudienceV2DeviceBrand, error) {
	ev := ToolsEstimateAudienceV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2DeviceBrand: valid values are %v", v, AllowedToolsEstimateAudienceV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_device_brand value
func (v ToolsEstimateAudienceV2DeviceBrand) Ptr() *ToolsEstimateAudienceV2DeviceBrand {
	return &v
}
