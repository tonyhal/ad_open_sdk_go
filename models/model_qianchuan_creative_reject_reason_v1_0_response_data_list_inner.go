/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeRejectReasonV10ResponseDataListInner struct for QianchuanCreativeRejectReasonV10ResponseDataListInner
type QianchuanCreativeRejectReasonV10ResponseDataListInner struct {
	//
	AuditRecords []*QianchuanCreativeRejectReasonV10ResponseDataListInnerAuditRecordsInner `json:"audit_records,omitempty"`
	//
	CreativeId *int64 `json:"creative_id,omitempty"`
}
