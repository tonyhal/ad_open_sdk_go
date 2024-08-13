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

// ReportVideoFrameGetV2Metrics
type ReportVideoFrameGetV2Metrics string

// List of report_video_frame_get_v2_metrics
const (
	CLICK_CNT_ReportVideoFrameGetV2Metrics     ReportVideoFrameGetV2Metrics = "click_cnt"
	DY_COMMENT_ReportVideoFrameGetV2Metrics    ReportVideoFrameGetV2Metrics = "dy_comment"
	DY_FOLLOW_ReportVideoFrameGetV2Metrics     ReportVideoFrameGetV2Metrics = "dy_follow"
	DY_LIKE_ReportVideoFrameGetV2Metrics       ReportVideoFrameGetV2Metrics = "dy_like"
	USER_LOSE_CNT_ReportVideoFrameGetV2Metrics ReportVideoFrameGetV2Metrics = "user_lose_cnt"
)

// All allowed values of ReportVideoFrameGetV2Metrics enum
var AllowedReportVideoFrameGetV2MetricsEnumValues = []ReportVideoFrameGetV2Metrics{
	"click_cnt",
	"dy_comment",
	"dy_follow",
	"dy_like",
	"user_lose_cnt",
}

// NewReportVideoFrameGetV2MetricsFromValue returns a pointer to a valid ReportVideoFrameGetV2Metrics
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportVideoFrameGetV2MetricsFromValue(v string) (*ReportVideoFrameGetV2Metrics, error) {
	ev := ReportVideoFrameGetV2Metrics(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportVideoFrameGetV2Metrics: valid values are %v", v, AllowedReportVideoFrameGetV2MetricsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportVideoFrameGetV2Metrics) IsValid() bool {
	for _, existing := range AllowedReportVideoFrameGetV2MetricsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_video_frame_get_v2_metrics value
func (v ReportVideoFrameGetV2Metrics) Ptr() *ReportVideoFrameGetV2Metrics {
	return &v
}
