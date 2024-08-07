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

// QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType
type QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType string

// List of qianchuan_aweme_auth_list_get_v1.0_data_authorization_infos_auth_type
const (
	ALL_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType              QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "ALL"
	AWEME_COOPERATOR_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "AWEME_COOPERATOR"
	OFFICIAL_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType         QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "OFFICIAL"
	SELF_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType             QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "SELF"
)

// All allowed values of QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType enum
var AllowedQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthTypeEnumValues = []QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType{
	"ALL",
	"AWEME_COOPERATOR",
	"OFFICIAL",
	"SELF",
}

// NewQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthTypeFromValue returns a pointer to a valid QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthTypeFromValue(v string) (*QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType, error) {
	ev := QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType: valid values are %v", v, AllowedQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_auth_list_get_v1.0_data_authorization_infos_auth_type value
func (v QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType) Ptr() *QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType {
	return &v
}
