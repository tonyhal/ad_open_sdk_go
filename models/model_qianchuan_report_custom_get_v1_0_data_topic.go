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

// QianchuanReportCustomGetV10DataTopic
type QianchuanReportCustomGetV10DataTopic string

// List of qianchuan_report_custom_get_v1.0_data_topic
const (
	ECP_BASIC_DATA_QianchuanReportCustomGetV10DataTopic                 QianchuanReportCustomGetV10DataTopic = "ECP_BASIC_DATA"
	SITE_PROMOTION_POST_DATA_LIVE_QianchuanReportCustomGetV10DataTopic  QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_LIVE"
	SITE_PROMOTION_POST_DATA_OTHER_QianchuanReportCustomGetV10DataTopic QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_OTHER"
	SITE_PROMOTION_POST_DATA_TITLE_QianchuanReportCustomGetV10DataTopic QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_TITLE"
	SITE_PROMOTION_POST_DATA_VIDEO_QianchuanReportCustomGetV10DataTopic QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_VIDEO"
)

// All allowed values of QianchuanReportCustomGetV10DataTopic enum
var AllowedQianchuanReportCustomGetV10DataTopicEnumValues = []QianchuanReportCustomGetV10DataTopic{
	"ECP_BASIC_DATA",
	"SITE_PROMOTION_POST_DATA_LIVE",
	"SITE_PROMOTION_POST_DATA_OTHER",
	"SITE_PROMOTION_POST_DATA_TITLE",
	"SITE_PROMOTION_POST_DATA_VIDEO",
}

// NewQianchuanReportCustomGetV10DataTopicFromValue returns a pointer to a valid QianchuanReportCustomGetV10DataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCustomGetV10DataTopicFromValue(v string) (*QianchuanReportCustomGetV10DataTopic, error) {
	ev := QianchuanReportCustomGetV10DataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCustomGetV10DataTopic: valid values are %v", v, AllowedQianchuanReportCustomGetV10DataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCustomGetV10DataTopic) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCustomGetV10DataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_custom_get_v1.0_data_topic value
func (v QianchuanReportCustomGetV10DataTopic) Ptr() *QianchuanReportCustomGetV10DataTopic {
	return &v
}
