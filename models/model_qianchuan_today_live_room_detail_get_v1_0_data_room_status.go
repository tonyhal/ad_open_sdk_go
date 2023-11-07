/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanTodayLiveRoomDetailGetV10DataRoomStatus
type QianchuanTodayLiveRoomDetailGetV10DataRoomStatus string

// List of qianchuan_today_live_room_detail_get_v1.0_data_room_status
const (
	ALL_QianchuanTodayLiveRoomDetailGetV10DataRoomStatus     QianchuanTodayLiveRoomDetailGetV10DataRoomStatus = "ALL"
	FINISH_QianchuanTodayLiveRoomDetailGetV10DataRoomStatus  QianchuanTodayLiveRoomDetailGetV10DataRoomStatus = "FINISH"
	LIVING_QianchuanTodayLiveRoomDetailGetV10DataRoomStatus  QianchuanTodayLiveRoomDetailGetV10DataRoomStatus = "LIVING"
	PAUSE_QianchuanTodayLiveRoomDetailGetV10DataRoomStatus   QianchuanTodayLiveRoomDetailGetV10DataRoomStatus = "PAUSE"
	PREPARE_QianchuanTodayLiveRoomDetailGetV10DataRoomStatus QianchuanTodayLiveRoomDetailGetV10DataRoomStatus = "PREPARE"
	UNKNOWN_QianchuanTodayLiveRoomDetailGetV10DataRoomStatus QianchuanTodayLiveRoomDetailGetV10DataRoomStatus = "UNKNOWN"
)

// All allowed values of QianchuanTodayLiveRoomDetailGetV10DataRoomStatus enum
var AllowedQianchuanTodayLiveRoomDetailGetV10DataRoomStatusEnumValues = []QianchuanTodayLiveRoomDetailGetV10DataRoomStatus{
	"ALL",
	"FINISH",
	"LIVING",
	"PAUSE",
	"PREPARE",
	"UNKNOWN",
}

// NewQianchuanTodayLiveRoomDetailGetV10DataRoomStatusFromValue returns a pointer to a valid QianchuanTodayLiveRoomDetailGetV10DataRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomDetailGetV10DataRoomStatusFromValue(v string) (*QianchuanTodayLiveRoomDetailGetV10DataRoomStatus, error) {
	ev := QianchuanTodayLiveRoomDetailGetV10DataRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomDetailGetV10DataRoomStatus: valid values are %v", v, AllowedQianchuanTodayLiveRoomDetailGetV10DataRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomDetailGetV10DataRoomStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomDetailGetV10DataRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_detail_get_v1.0_data_room_status value
func (v QianchuanTodayLiveRoomDetailGetV10DataRoomStatus) Ptr() *QianchuanTodayLiveRoomDetailGetV10DataRoomStatus {
	return &v
}
