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

// ReportLiveRoomAudiencePortraitGetV2Dimension
type ReportLiveRoomAudiencePortraitGetV2Dimension string

// List of report_live_room_audience_portrait_get_v2_dimension
const (
	PROVINCE_ReportLiveRoomAudiencePortraitGetV2Dimension ReportLiveRoomAudiencePortraitGetV2Dimension = "PROVINCE"
	CITY_ReportLiveRoomAudiencePortraitGetV2Dimension     ReportLiveRoomAudiencePortraitGetV2Dimension = "CITY"
	GENDER_ReportLiveRoomAudiencePortraitGetV2Dimension   ReportLiveRoomAudiencePortraitGetV2Dimension = "GENDER"
	PLATFORM_ReportLiveRoomAudiencePortraitGetV2Dimension ReportLiveRoomAudiencePortraitGetV2Dimension = "PLATFORM"
	AGE_ReportLiveRoomAudiencePortraitGetV2Dimension      ReportLiveRoomAudiencePortraitGetV2Dimension = "AGE"
)

// All allowed values of ReportLiveRoomAudiencePortraitGetV2Dimension enum
var AllowedReportLiveRoomAudiencePortraitGetV2DimensionEnumValues = []ReportLiveRoomAudiencePortraitGetV2Dimension{
	"PROVINCE",
	"CITY",
	"GENDER",
	"PLATFORM",
	"AGE",
}

// NewReportLiveRoomAudiencePortraitGetV2DimensionFromValue returns a pointer to a valid ReportLiveRoomAudiencePortraitGetV2Dimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomAudiencePortraitGetV2DimensionFromValue(v string) (*ReportLiveRoomAudiencePortraitGetV2Dimension, error) {
	ev := ReportLiveRoomAudiencePortraitGetV2Dimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomAudiencePortraitGetV2Dimension: valid values are %v", v, AllowedReportLiveRoomAudiencePortraitGetV2DimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomAudiencePortraitGetV2Dimension) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomAudiencePortraitGetV2DimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_audience_portrait_get_v2_dimension value
func (v ReportLiveRoomAudiencePortraitGetV2Dimension) Ptr() *ReportLiveRoomAudiencePortraitGetV2Dimension {
	return &v
}
