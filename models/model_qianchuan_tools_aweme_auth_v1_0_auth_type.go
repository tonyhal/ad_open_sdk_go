/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsAwemeAuthV10AuthType
type QianchuanToolsAwemeAuthV10AuthType string

// List of qianchuan_tools_aweme_auth_v1.0_auth_type
const (
	SELF_QianchuanToolsAwemeAuthV10AuthType             QianchuanToolsAwemeAuthV10AuthType = "SELF"
	AWEME_COOPERATOR_QianchuanToolsAwemeAuthV10AuthType QianchuanToolsAwemeAuthV10AuthType = "AWEME_COOPERATOR"
)

// All allowed values of QianchuanToolsAwemeAuthV10AuthType enum
var AllowedQianchuanToolsAwemeAuthV10AuthTypeEnumValues = []QianchuanToolsAwemeAuthV10AuthType{
	"SELF",
	"AWEME_COOPERATOR",
}

// NewQianchuanToolsAwemeAuthV10AuthTypeFromValue returns a pointer to a valid QianchuanToolsAwemeAuthV10AuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsAwemeAuthV10AuthTypeFromValue(v string) (*QianchuanToolsAwemeAuthV10AuthType, error) {
	ev := QianchuanToolsAwemeAuthV10AuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsAwemeAuthV10AuthType: valid values are %v", v, AllowedQianchuanToolsAwemeAuthV10AuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsAwemeAuthV10AuthType) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsAwemeAuthV10AuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_aweme_auth_v1.0_auth_type value
func (v QianchuanToolsAwemeAuthV10AuthType) Ptr() *QianchuanToolsAwemeAuthV10AuthType {
	return &v
}
