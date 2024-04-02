/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListAdDownloadStatus
type PromotionListV30DataListAdDownloadStatus string

// List of promotion_list_v3.0_data_list_ad_download_status
const (
	OFF_PromotionListV30DataListAdDownloadStatus PromotionListV30DataListAdDownloadStatus = "OFF"
	ON_PromotionListV30DataListAdDownloadStatus  PromotionListV30DataListAdDownloadStatus = "ON"
)

// All allowed values of PromotionListV30DataListAdDownloadStatus enum
var AllowedPromotionListV30DataListAdDownloadStatusEnumValues = []PromotionListV30DataListAdDownloadStatus{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListAdDownloadStatusFromValue returns a pointer to a valid PromotionListV30DataListAdDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListAdDownloadStatusFromValue(v string) (*PromotionListV30DataListAdDownloadStatus, error) {
	ev := PromotionListV30DataListAdDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListAdDownloadStatus: valid values are %v", v, AllowedPromotionListV30DataListAdDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListAdDownloadStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListAdDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_ad_download_status value
func (v PromotionListV30DataListAdDownloadStatus) Ptr() *PromotionListV30DataListAdDownloadStatus {
	return &v
}
