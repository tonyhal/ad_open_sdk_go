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

// QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform
type QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform string

// List of qianchuan_creative_reject_reason_v1.0_data_list_audit_records_audit_platform
const (
	CONTENT_QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform = "CONTENT"
	AD_QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform      QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform = "AD"
	UNKNOWN_QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform = "UNKNOWN"
)

// All allowed values of QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform enum
var AllowedQianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatformEnumValues = []QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform{
	"CONTENT",
	"AD",
	"UNKNOWN",
}

// NewQianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatformFromValue returns a pointer to a valid QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatformFromValue(v string) (*QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform, error) {
	ev := QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform: valid values are %v", v, AllowedQianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_reject_reason_v1.0_data_list_audit_records_audit_platform value
func (v QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform) Ptr() *QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform {
	return &v
}