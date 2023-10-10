/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanTodayLiveRoomUserGetV10Dimension
type QianchuanTodayLiveRoomUserGetV10Dimension string

// List of qianchuan_today_live_room_user_get_v1.0_dimension
const (
	AGE_QianchuanTodayLiveRoomUserGetV10Dimension    QianchuanTodayLiveRoomUserGetV10Dimension = "AGE"
	CITY_QianchuanTodayLiveRoomUserGetV10Dimension   QianchuanTodayLiveRoomUserGetV10Dimension = "CITY"
	GENDER_QianchuanTodayLiveRoomUserGetV10Dimension QianchuanTodayLiveRoomUserGetV10Dimension = "GENDER"
)

// All allowed values of QianchuanTodayLiveRoomUserGetV10Dimension enum
var AllowedQianchuanTodayLiveRoomUserGetV10DimensionEnumValues = []QianchuanTodayLiveRoomUserGetV10Dimension{
	"AGE",
	"CITY",
	"GENDER",
}

// NewQianchuanTodayLiveRoomUserGetV10DimensionFromValue returns a pointer to a valid QianchuanTodayLiveRoomUserGetV10Dimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomUserGetV10DimensionFromValue(v string) (*QianchuanTodayLiveRoomUserGetV10Dimension, error) {
	ev := QianchuanTodayLiveRoomUserGetV10Dimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomUserGetV10Dimension: valid values are %v", v, AllowedQianchuanTodayLiveRoomUserGetV10DimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomUserGetV10Dimension) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomUserGetV10DimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_user_get_v1.0_dimension value
func (v QianchuanTodayLiveRoomUserGetV10Dimension) Ptr() *QianchuanTodayLiveRoomUserGetV10Dimension {
	return &v
}
