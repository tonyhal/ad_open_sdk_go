/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaAssetsDetailReadV2DataListStatus
type DpaAssetsDetailReadV2DataListStatus string

// List of dpa_assets_detail_read_v2_data_list_status
const (
	ENABLE_DpaAssetsDetailReadV2DataListStatus  DpaAssetsDetailReadV2DataListStatus = "ENABLE"
	DISABLE_DpaAssetsDetailReadV2DataListStatus DpaAssetsDetailReadV2DataListStatus = "DISABLE"
)

// All allowed values of DpaAssetsDetailReadV2DataListStatus enum
var AllowedDpaAssetsDetailReadV2DataListStatusEnumValues = []DpaAssetsDetailReadV2DataListStatus{
	"ENABLE",
	"DISABLE",
}

// NewDpaAssetsDetailReadV2DataListStatusFromValue returns a pointer to a valid DpaAssetsDetailReadV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsDetailReadV2DataListStatusFromValue(v string) (*DpaAssetsDetailReadV2DataListStatus, error) {
	ev := DpaAssetsDetailReadV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsDetailReadV2DataListStatus: valid values are %v", v, AllowedDpaAssetsDetailReadV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsDetailReadV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetsDetailReadV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_detail_read_v2_data_list_status value
func (v DpaAssetsDetailReadV2DataListStatus) Ptr() *DpaAssetsDetailReadV2DataListStatus {
	return &v
}
