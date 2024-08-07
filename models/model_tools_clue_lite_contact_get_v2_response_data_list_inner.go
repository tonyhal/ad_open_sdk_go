/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLiteContactGetV2ResponseDataListInner struct for ToolsClueLiteContactGetV2ResponseDataListInner
type ToolsClueLiteContactGetV2ResponseDataListInner struct {
	CallDirection *ToolsClueLiteContactGetV2DataListCallDirection `json:"call_direction,omitempty"`
	//
	CallDuration *int64 `json:"call_duration,omitempty"`
	//
	CalleeNumber *string `json:"callee_number,omitempty"`
	//
	CallerNumber *string `json:"caller_number,omitempty"`
	//
	ClueId *int64 `json:"clue_id,omitempty"`
	//
	ContactId *string `json:"contact_id,omitempty"`
	//
	Duration *int64 `json:"duration,omitempty"`
	//
	EndStateShowCode *int64 `json:"end_state_show_code,omitempty"`
	//
	EndStateShowMsg *string `json:"end_state_show_msg,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
	//
	VirtualNumber *string `json:"virtual_number,omitempty"`
}
