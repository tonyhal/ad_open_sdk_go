/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceDeviceBrand
type ProjectUpdateV30AudienceDeviceBrand string

// List of project_update_v3.0_audience_device_brand
const (
	Enum_360_ProjectUpdateV30AudienceDeviceBrand    ProjectUpdateV30AudienceDeviceBrand = "360"
	APPLE_ProjectUpdateV30AudienceDeviceBrand       ProjectUpdateV30AudienceDeviceBrand = "APPLE"
	CHINAMOBILE_ProjectUpdateV30AudienceDeviceBrand ProjectUpdateV30AudienceDeviceBrand = "CHINAMOBILE"
	COOLPAD_ProjectUpdateV30AudienceDeviceBrand     ProjectUpdateV30AudienceDeviceBrand = "COOLPAD"
	GIONEE_ProjectUpdateV30AudienceDeviceBrand      ProjectUpdateV30AudienceDeviceBrand = "GIONEE"
	GOOGLE_ProjectUpdateV30AudienceDeviceBrand      ProjectUpdateV30AudienceDeviceBrand = "GOOGLE"
	HISENSE_ProjectUpdateV30AudienceDeviceBrand     ProjectUpdateV30AudienceDeviceBrand = "HISENSE"
	HONOR_ProjectUpdateV30AudienceDeviceBrand       ProjectUpdateV30AudienceDeviceBrand = "HONOR"
	HTC_ProjectUpdateV30AudienceDeviceBrand         ProjectUpdateV30AudienceDeviceBrand = "HTC"
	HUAWEI_ProjectUpdateV30AudienceDeviceBrand      ProjectUpdateV30AudienceDeviceBrand = "HUAWEI"
	LENOVO_ProjectUpdateV30AudienceDeviceBrand      ProjectUpdateV30AudienceDeviceBrand = "LENOVO"
	LETV_ProjectUpdateV30AudienceDeviceBrand        ProjectUpdateV30AudienceDeviceBrand = "LETV"
	LG_ProjectUpdateV30AudienceDeviceBrand          ProjectUpdateV30AudienceDeviceBrand = "LG"
	MEIZU_ProjectUpdateV30AudienceDeviceBrand       ProjectUpdateV30AudienceDeviceBrand = "MEIZU"
	MOTO_ProjectUpdateV30AudienceDeviceBrand        ProjectUpdateV30AudienceDeviceBrand = "MOTO"
	NOKIA_ProjectUpdateV30AudienceDeviceBrand       ProjectUpdateV30AudienceDeviceBrand = "NOKIA"
	NUBIA_ProjectUpdateV30AudienceDeviceBrand       ProjectUpdateV30AudienceDeviceBrand = "NUBIA"
	ONEPLUS_ProjectUpdateV30AudienceDeviceBrand     ProjectUpdateV30AudienceDeviceBrand = "ONEPLUS"
	OPPO_ProjectUpdateV30AudienceDeviceBrand        ProjectUpdateV30AudienceDeviceBrand = "OPPO"
	PEPPER_ProjectUpdateV30AudienceDeviceBrand      ProjectUpdateV30AudienceDeviceBrand = "PEPPER"
	QIKU_ProjectUpdateV30AudienceDeviceBrand        ProjectUpdateV30AudienceDeviceBrand = "QIKU"
	SAMSUNG_ProjectUpdateV30AudienceDeviceBrand     ProjectUpdateV30AudienceDeviceBrand = "SAMSUNG"
	SMARTISAN_ProjectUpdateV30AudienceDeviceBrand   ProjectUpdateV30AudienceDeviceBrand = "SMARTISAN"
	SONY_ProjectUpdateV30AudienceDeviceBrand        ProjectUpdateV30AudienceDeviceBrand = "SONY"
	TCL_ProjectUpdateV30AudienceDeviceBrand         ProjectUpdateV30AudienceDeviceBrand = "TCL"
	VIVO_ProjectUpdateV30AudienceDeviceBrand        ProjectUpdateV30AudienceDeviceBrand = "VIVO"
	XIAOMI_ProjectUpdateV30AudienceDeviceBrand      ProjectUpdateV30AudienceDeviceBrand = "XIAOMI"
	ZTE_ProjectUpdateV30AudienceDeviceBrand         ProjectUpdateV30AudienceDeviceBrand = "ZTE"
)

// All allowed values of ProjectUpdateV30AudienceDeviceBrand enum
var AllowedProjectUpdateV30AudienceDeviceBrandEnumValues = []ProjectUpdateV30AudienceDeviceBrand{
	"360",
	"APPLE",
	"CHINAMOBILE",
	"COOLPAD",
	"GIONEE",
	"GOOGLE",
	"HISENSE",
	"HONOR",
	"HTC",
	"HUAWEI",
	"LENOVO",
	"LETV",
	"LG",
	"MEIZU",
	"MOTO",
	"NOKIA",
	"NUBIA",
	"ONEPLUS",
	"OPPO",
	"PEPPER",
	"QIKU",
	"SAMSUNG",
	"SMARTISAN",
	"SONY",
	"TCL",
	"VIVO",
	"XIAOMI",
	"ZTE",
}

// NewProjectUpdateV30AudienceDeviceBrandFromValue returns a pointer to a valid ProjectUpdateV30AudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceDeviceBrandFromValue(v string) (*ProjectUpdateV30AudienceDeviceBrand, error) {
	ev := ProjectUpdateV30AudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceDeviceBrand: valid values are %v", v, AllowedProjectUpdateV30AudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_device_brand value
func (v ProjectUpdateV30AudienceDeviceBrand) Ptr() *ProjectUpdateV30AudienceDeviceBrand {
	return &v
}
