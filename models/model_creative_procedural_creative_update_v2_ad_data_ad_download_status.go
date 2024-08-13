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

// CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus
type CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus int64

// List of creative_procedural_creative_update_v2_ad_data_ad_download_status
const (
	Enum_0_CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus = 0
	Enum_1_CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus = 1
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataAdDownloadStatusEnumValues = []CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus{
	0,
	1,
}

// NewCreativeProceduralCreativeUpdateV2AdDataAdDownloadStatusFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataAdDownloadStatusFromValue(v int64) (*CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataAdDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataAdDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_ad_download_status value
func (v CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus) Ptr() *CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus {
	return &v
}
