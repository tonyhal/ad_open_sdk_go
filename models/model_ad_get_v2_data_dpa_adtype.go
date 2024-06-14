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

// AdGetV2DataDpaAdtype
type AdGetV2DataDpaAdtype string

// List of ad_get_v2_data_dpa_adtype
const (
	DPA_LINK_AdGetV2DataDpaAdtype AdGetV2DataDpaAdtype = "DPA_LINK"
	DPA_APP_AdGetV2DataDpaAdtype  AdGetV2DataDpaAdtype = "DPA_APP"
)

// All allowed values of AdGetV2DataDpaAdtype enum
var AllowedAdGetV2DataDpaAdtypeEnumValues = []AdGetV2DataDpaAdtype{
	"DPA_LINK",
	"DPA_APP",
}

// NewAdGetV2DataDpaAdtypeFromValue returns a pointer to a valid AdGetV2DataDpaAdtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDpaAdtypeFromValue(v string) (*AdGetV2DataDpaAdtype, error) {
	ev := AdGetV2DataDpaAdtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDpaAdtype: valid values are %v", v, AllowedAdGetV2DataDpaAdtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDpaAdtype) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDpaAdtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_dpa_adtype value
func (v AdGetV2DataDpaAdtype) Ptr() *AdGetV2DataDpaAdtype {
	return &v
}
