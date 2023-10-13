/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10DeliverySettingQcpxMode
type QianchuanAdUpdateV10DeliverySettingQcpxMode string

// List of qianchuan_ad_update_v1.0_delivery_setting_qcpx_mode
const (
	QCPX_MODE_DEFAULT_QianchuanAdUpdateV10DeliverySettingQcpxMode QianchuanAdUpdateV10DeliverySettingQcpxMode = "QCPX_MODE_DEFAULT"
	QCPX_MODE_OFF_QianchuanAdUpdateV10DeliverySettingQcpxMode     QianchuanAdUpdateV10DeliverySettingQcpxMode = "QCPX_MODE_OFF"
	QCPX_MODE_ON_QianchuanAdUpdateV10DeliverySettingQcpxMode      QianchuanAdUpdateV10DeliverySettingQcpxMode = "QCPX_MODE_ON"
)

// All allowed values of QianchuanAdUpdateV10DeliverySettingQcpxMode enum
var AllowedQianchuanAdUpdateV10DeliverySettingQcpxModeEnumValues = []QianchuanAdUpdateV10DeliverySettingQcpxMode{
	"QCPX_MODE_DEFAULT",
	"QCPX_MODE_OFF",
	"QCPX_MODE_ON",
}

// NewQianchuanAdUpdateV10DeliverySettingQcpxModeFromValue returns a pointer to a valid QianchuanAdUpdateV10DeliverySettingQcpxMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10DeliverySettingQcpxModeFromValue(v string) (*QianchuanAdUpdateV10DeliverySettingQcpxMode, error) {
	ev := QianchuanAdUpdateV10DeliverySettingQcpxMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10DeliverySettingQcpxMode: valid values are %v", v, AllowedQianchuanAdUpdateV10DeliverySettingQcpxModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10DeliverySettingQcpxMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10DeliverySettingQcpxModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_delivery_setting_qcpx_mode value
func (v QianchuanAdUpdateV10DeliverySettingQcpxMode) Ptr() *QianchuanAdUpdateV10DeliverySettingQcpxMode {
	return &v
}
