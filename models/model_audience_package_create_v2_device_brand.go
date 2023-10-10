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

// AudiencePackageCreateV2DeviceBrand
type AudiencePackageCreateV2DeviceBrand string

// List of audience_package_create_v2_device_brand
const (
	NUBIA_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "NUBIA"
	XIAOMI_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "XIAOMI"
	VIVO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "VIVO"
	GOOGLE_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "GOOGLE"
	LETV_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "LETV"
	HTC_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "HTC"
	ZTE_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "ZTE"
	ONEPLUS_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "ONEPLUS"
	SMARTISAN_AudiencePackageCreateV2DeviceBrand   AudiencePackageCreateV2DeviceBrand = "SMARTISAN"
	Enum_360_AudiencePackageCreateV2DeviceBrand    AudiencePackageCreateV2DeviceBrand = "360"
	MOTO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "MOTO"
	SAMSUNG_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "SAMSUNG"
	LG_AudiencePackageCreateV2DeviceBrand          AudiencePackageCreateV2DeviceBrand = "LG"
	TCL_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "TCL"
	APPLE_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "APPLE"
	HONOR_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "HONOR"
	NOKIA_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "NOKIA"
	CHINAMOBILE_AudiencePackageCreateV2DeviceBrand AudiencePackageCreateV2DeviceBrand = "CHINAMOBILE"
	MEIZU_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "MEIZU"
	SONY_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "SONY"
	PEPPER_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "PEPPER"
	HISENSE_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "HISENSE"
	OPPO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "OPPO"
	HUAWEI_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "HUAWEI"
	COOLPAD_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "COOLPAD"
	LENOVO_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "LENOVO"
	GIONEE_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "GIONEE"
	QIKU_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "QIKU"
)

// All allowed values of AudiencePackageCreateV2DeviceBrand enum
var AllowedAudiencePackageCreateV2DeviceBrandEnumValues = []AudiencePackageCreateV2DeviceBrand{
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

// NewAudiencePackageCreateV2DeviceBrandFromValue returns a pointer to a valid AudiencePackageCreateV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2DeviceBrandFromValue(v string) (*AudiencePackageCreateV2DeviceBrand, error) {
	ev := AudiencePackageCreateV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2DeviceBrand: valid values are %v", v, AllowedAudiencePackageCreateV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_device_brand value
func (v AudiencePackageCreateV2DeviceBrand) Ptr() *AudiencePackageCreateV2DeviceBrand {
	return &v
}
