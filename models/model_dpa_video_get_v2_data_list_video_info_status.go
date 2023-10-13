/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaVideoGetV2DataListVideoInfoStatus
type DpaVideoGetV2DataListVideoInfoStatus string

// List of dpa_video_get_v2_data_list_video_info_status
const (
	AVAILABLE_DpaVideoGetV2DataListVideoInfoStatus   DpaVideoGetV2DataListVideoInfoStatus = "AVAILABLE"
	UNAVAILABLE_DpaVideoGetV2DataListVideoInfoStatus DpaVideoGetV2DataListVideoInfoStatus = "UNAVAILABLE"
)

// All allowed values of DpaVideoGetV2DataListVideoInfoStatus enum
var AllowedDpaVideoGetV2DataListVideoInfoStatusEnumValues = []DpaVideoGetV2DataListVideoInfoStatus{
	"AVAILABLE",
	"UNAVAILABLE",
}

// NewDpaVideoGetV2DataListVideoInfoStatusFromValue returns a pointer to a valid DpaVideoGetV2DataListVideoInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaVideoGetV2DataListVideoInfoStatusFromValue(v string) (*DpaVideoGetV2DataListVideoInfoStatus, error) {
	ev := DpaVideoGetV2DataListVideoInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaVideoGetV2DataListVideoInfoStatus: valid values are %v", v, AllowedDpaVideoGetV2DataListVideoInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaVideoGetV2DataListVideoInfoStatus) IsValid() bool {
	for _, existing := range AllowedDpaVideoGetV2DataListVideoInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_video_get_v2_data_list_video_info_status value
func (v DpaVideoGetV2DataListVideoInfoStatus) Ptr() *DpaVideoGetV2DataListVideoInfoStatus {
	return &v
}
