/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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
	HONOR_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "HONOR"
	CHINAMOBILE_ToolsEstimateAudienceV2DeviceBrand ToolsEstimateAudienceV2DeviceBrand = "CHINAMOBILE"
	MOTO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "MOTO"
	NUBIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NUBIA"
	APPLE_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "APPLE"
	GIONEE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GIONEE"
	HUAWEI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "HUAWEI"
	SONY_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "SONY"
	HISENSE_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "HISENSE"
	QIKU_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "QIKU"
	SMARTISAN_ToolsEstimateAudienceV2DeviceBrand   ToolsEstimateAudienceV2DeviceBrand = "SMARTISAN"
	XIAOMI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "XIAOMI"
	LETV_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "LETV"
	ZTE_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "ZTE"
	Enum_360_ToolsEstimateAudienceV2DeviceBrand    ToolsEstimateAudienceV2DeviceBrand = "360"
	COOLPAD_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "COOLPAD"
	SAMSUNG_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "SAMSUNG"
	LG_ToolsEstimateAudienceV2DeviceBrand          ToolsEstimateAudienceV2DeviceBrand = "LG"
	PEPPER_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "PEPPER"
	LENOVO_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "LENOVO"
	MEIZU_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "MEIZU"
	ONEPLUS_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "ONEPLUS"
	OPPO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "OPPO"
	VIVO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "VIVO"
	NOKIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NOKIA"
	HTC_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "HTC"
	TCL_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "TCL"
	GOOGLE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GOOGLE"
)

// All allowed values of ToolsEstimateAudienceV2DeviceBrand enum
var AllowedToolsEstimateAudienceV2DeviceBrandEnumValues = []ToolsEstimateAudienceV2DeviceBrand{
	"HONOR",
	"CHINAMOBILE",
	"MOTO",
	"NUBIA",
	"APPLE",
	"GIONEE",
	"HUAWEI",
	"SONY",
	"HISENSE",
	"QIKU",
	"SMARTISAN",
	"XIAOMI",
	"LETV",
	"ZTE",
	"360",
	"COOLPAD",
	"SAMSUNG",
	"LG",
	"PEPPER",
	"LENOVO",
	"MEIZU",
	"ONEPLUS",
	"OPPO",
	"VIVO",
	"NOKIA",
	"HTC",
	"TCL",
	"GOOGLE",
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
