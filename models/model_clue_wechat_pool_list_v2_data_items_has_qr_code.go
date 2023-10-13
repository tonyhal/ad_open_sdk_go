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

// ClueWechatPoolListV2DataItemsHasQrCode
type ClueWechatPoolListV2DataItemsHasQrCode string

// List of clue_wechat_pool_list_v2_data_items_has_qr_code
const (
	TRUE_ClueWechatPoolListV2DataItemsHasQrCode  ClueWechatPoolListV2DataItemsHasQrCode = "TRUE"
	FALSE_ClueWechatPoolListV2DataItemsHasQrCode ClueWechatPoolListV2DataItemsHasQrCode = "FALSE"
)

// All allowed values of ClueWechatPoolListV2DataItemsHasQrCode enum
var AllowedClueWechatPoolListV2DataItemsHasQrCodeEnumValues = []ClueWechatPoolListV2DataItemsHasQrCode{
	"TRUE",
	"FALSE",
}

// NewClueWechatPoolListV2DataItemsHasQrCodeFromValue returns a pointer to a valid ClueWechatPoolListV2DataItemsHasQrCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatPoolListV2DataItemsHasQrCodeFromValue(v string) (*ClueWechatPoolListV2DataItemsHasQrCode, error) {
	ev := ClueWechatPoolListV2DataItemsHasQrCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatPoolListV2DataItemsHasQrCode: valid values are %v", v, AllowedClueWechatPoolListV2DataItemsHasQrCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatPoolListV2DataItemsHasQrCode) IsValid() bool {
	for _, existing := range AllowedClueWechatPoolListV2DataItemsHasQrCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_pool_list_v2_data_items_has_qr_code value
func (v ClueWechatPoolListV2DataItemsHasQrCode) Ptr() *ClueWechatPoolListV2DataItemsHasQrCode {
	return &v
}
