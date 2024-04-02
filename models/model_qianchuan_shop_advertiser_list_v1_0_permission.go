/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanShopAdvertiserListV10Permission
type QianchuanShopAdvertiserListV10Permission string

// List of qianchuan_shop_advertiser_list_v1.0_permission
const (
	QC_AWEME_QianchuanShopAdvertiserListV10Permission QianchuanShopAdvertiserListV10Permission = "QC_AWEME"
)

// All allowed values of QianchuanShopAdvertiserListV10Permission enum
var AllowedQianchuanShopAdvertiserListV10PermissionEnumValues = []QianchuanShopAdvertiserListV10Permission{
	"QC_AWEME",
}

// NewQianchuanShopAdvertiserListV10PermissionFromValue returns a pointer to a valid QianchuanShopAdvertiserListV10Permission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanShopAdvertiserListV10PermissionFromValue(v string) (*QianchuanShopAdvertiserListV10Permission, error) {
	ev := QianchuanShopAdvertiserListV10Permission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanShopAdvertiserListV10Permission: valid values are %v", v, AllowedQianchuanShopAdvertiserListV10PermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanShopAdvertiserListV10Permission) IsValid() bool {
	for _, existing := range AllowedQianchuanShopAdvertiserListV10PermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_shop_advertiser_list_v1.0_permission value
func (v QianchuanShopAdvertiserListV10Permission) Ptr() *QianchuanShopAdvertiserListV10Permission {
	return &v
}
