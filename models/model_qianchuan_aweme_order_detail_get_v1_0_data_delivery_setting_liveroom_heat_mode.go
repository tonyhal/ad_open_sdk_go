/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode
type QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode string

// List of qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_liveroom_heat_mode
const (
	BY_ROOM_QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode  QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode = "BY_ROOM"
	BY_VIDEO_QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode = "BY_VIDEO"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode enum
var AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatModeEnumValues = []QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode{
	"BY_ROOM",
	"BY_VIDEO",
}

// NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatModeFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatModeFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_liveroom_heat_mode value
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode) Ptr() *QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode {
	return &v
}
