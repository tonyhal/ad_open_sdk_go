/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRejectReasonV10ResponseDataListInner struct for QianchuanAdRejectReasonV10ResponseDataListInner
type QianchuanAdRejectReasonV10ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AuditRecords []*QianchuanAdRejectReasonV10ResponseDataListInnerAuditRecordsInner `json:"audit_records,omitempty"`
}
