/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueWechatInstanceDetailV2DataWechatListHasQrCode
type ClueWechatInstanceDetailV2DataWechatListHasQrCode string

// List of clue_wechat_instance_detail_v2_data_wechat_list_has_qr_code
const (
	FALSE_ClueWechatInstanceDetailV2DataWechatListHasQrCode ClueWechatInstanceDetailV2DataWechatListHasQrCode = "FALSE"
	TRUE_ClueWechatInstanceDetailV2DataWechatListHasQrCode  ClueWechatInstanceDetailV2DataWechatListHasQrCode = "TRUE"
)

// All allowed values of ClueWechatInstanceDetailV2DataWechatListHasQrCode enum
var AllowedClueWechatInstanceDetailV2DataWechatListHasQrCodeEnumValues = []ClueWechatInstanceDetailV2DataWechatListHasQrCode{
	"FALSE",
	"TRUE",
}

// NewClueWechatInstanceDetailV2DataWechatListHasQrCodeFromValue returns a pointer to a valid ClueWechatInstanceDetailV2DataWechatListHasQrCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatInstanceDetailV2DataWechatListHasQrCodeFromValue(v string) (*ClueWechatInstanceDetailV2DataWechatListHasQrCode, error) {
	ev := ClueWechatInstanceDetailV2DataWechatListHasQrCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatInstanceDetailV2DataWechatListHasQrCode: valid values are %v", v, AllowedClueWechatInstanceDetailV2DataWechatListHasQrCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatInstanceDetailV2DataWechatListHasQrCode) IsValid() bool {
	for _, existing := range AllowedClueWechatInstanceDetailV2DataWechatListHasQrCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_instance_detail_v2_data_wechat_list_has_qr_code value
func (v ClueWechatInstanceDetailV2DataWechatListHasQrCode) Ptr() *ClueWechatInstanceDetailV2DataWechatListHasQrCode {
	return &v
}
