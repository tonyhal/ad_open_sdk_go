/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic
type QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic string

// List of qianchuan_qianchuan_report_ltoday_live_room_data_get_v1.0_data_topic
const (
	ROOM_FLOW_PERFORMANCE_QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic = "ROOM_FLOW_PERFORMANCE"
	ROOM_PRODUCT_LIST_QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic     QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic = "ROOM_PRODUCT_LIST"
)

// All allowed values of QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic enum
var AllowedQianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopicEnumValues = []QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic{
	"ROOM_FLOW_PERFORMANCE",
	"ROOM_PRODUCT_LIST",
}

// NewQianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopicFromValue returns a pointer to a valid QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopicFromValue(v string) (*QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic, error) {
	ev := QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic: valid values are %v", v, AllowedQianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic) IsValid() bool {
	for _, existing := range AllowedQianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_qianchuan_report_ltoday_live_room_data_get_v1.0_data_topic value
func (v QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic) Ptr() *QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic {
	return &v
}
