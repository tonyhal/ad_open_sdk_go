/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanTodayLiveRoomGetV10AdStatus
type QianchuanTodayLiveRoomGetV10AdStatus string

// List of qianchuan_today_live_room_get_v1.0_ad_status
const (
	ALL_QianchuanTodayLiveRoomGetV10AdStatus         QianchuanTodayLiveRoomGetV10AdStatus = "ALL"
	DELIVERY_OK_QianchuanTodayLiveRoomGetV10AdStatus QianchuanTodayLiveRoomGetV10AdStatus = "DELIVERY_OK"
	NO_DELIVERY_QianchuanTodayLiveRoomGetV10AdStatus QianchuanTodayLiveRoomGetV10AdStatus = "NO_DELIVERY"
)

// All allowed values of QianchuanTodayLiveRoomGetV10AdStatus enum
var AllowedQianchuanTodayLiveRoomGetV10AdStatusEnumValues = []QianchuanTodayLiveRoomGetV10AdStatus{
	"ALL",
	"DELIVERY_OK",
	"NO_DELIVERY",
}

// NewQianchuanTodayLiveRoomGetV10AdStatusFromValue returns a pointer to a valid QianchuanTodayLiveRoomGetV10AdStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomGetV10AdStatusFromValue(v string) (*QianchuanTodayLiveRoomGetV10AdStatus, error) {
	ev := QianchuanTodayLiveRoomGetV10AdStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomGetV10AdStatus: valid values are %v", v, AllowedQianchuanTodayLiveRoomGetV10AdStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomGetV10AdStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomGetV10AdStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_get_v1.0_ad_status value
func (v QianchuanTodayLiveRoomGetV10AdStatus) Ptr() *QianchuanTodayLiveRoomGetV10AdStatus {
	return &v
}
