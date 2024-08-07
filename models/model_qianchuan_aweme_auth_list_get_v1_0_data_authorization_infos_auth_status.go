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

// QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus
type QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus string

// List of qianchuan_aweme_auth_list_get_v1.0_data_authorization_infos_auth_status
const (
	ALL_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus       QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus = "ALL"
	EFFECTIVE_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus = "EFFECTIVE"
	EXPIRED_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus   QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus = "EXPIRED"
	WAIT_BIND_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus = "WAIT_BIND"
)

// All allowed values of QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus enum
var AllowedQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatusEnumValues = []QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus{
	"ALL",
	"EFFECTIVE",
	"EXPIRED",
	"WAIT_BIND",
}

// NewQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatusFromValue returns a pointer to a valid QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatusFromValue(v string) (*QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus, error) {
	ev := QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus: valid values are %v", v, AllowedQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_auth_list_get_v1.0_data_authorization_infos_auth_status value
func (v QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus) Ptr() *QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus {
	return &v
}
