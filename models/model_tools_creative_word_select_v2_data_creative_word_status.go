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

// ToolsCreativeWordSelectV2DataCreativeWordStatus
type ToolsCreativeWordSelectV2DataCreativeWordStatus string

// List of tools_creative_word_select_v2_data_creative_word_status
const (
	CREATIVE_WORD_STATUS_APPROVE_ToolsCreativeWordSelectV2DataCreativeWordStatus ToolsCreativeWordSelectV2DataCreativeWordStatus = "CREATIVE_WORD_STATUS_APPROVE"
	CREATIVE_WORD_STATUS_AUDIT_ToolsCreativeWordSelectV2DataCreativeWordStatus   ToolsCreativeWordSelectV2DataCreativeWordStatus = "CREATIVE_WORD_STATUS_AUDIT"
	CREATIVE_WORD_STATUS_REAUDIT_ToolsCreativeWordSelectV2DataCreativeWordStatus ToolsCreativeWordSelectV2DataCreativeWordStatus = "CREATIVE_WORD_STATUS_REAUDIT"
	CREATIVE_WORD_STATUS_REJECT_ToolsCreativeWordSelectV2DataCreativeWordStatus  ToolsCreativeWordSelectV2DataCreativeWordStatus = "CREATIVE_WORD_STATUS_REJECT"
)

// All allowed values of ToolsCreativeWordSelectV2DataCreativeWordStatus enum
var AllowedToolsCreativeWordSelectV2DataCreativeWordStatusEnumValues = []ToolsCreativeWordSelectV2DataCreativeWordStatus{
	"CREATIVE_WORD_STATUS_APPROVE",
	"CREATIVE_WORD_STATUS_AUDIT",
	"CREATIVE_WORD_STATUS_REAUDIT",
	"CREATIVE_WORD_STATUS_REJECT",
}

// NewToolsCreativeWordSelectV2DataCreativeWordStatusFromValue returns a pointer to a valid ToolsCreativeWordSelectV2DataCreativeWordStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCreativeWordSelectV2DataCreativeWordStatusFromValue(v string) (*ToolsCreativeWordSelectV2DataCreativeWordStatus, error) {
	ev := ToolsCreativeWordSelectV2DataCreativeWordStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCreativeWordSelectV2DataCreativeWordStatus: valid values are %v", v, AllowedToolsCreativeWordSelectV2DataCreativeWordStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCreativeWordSelectV2DataCreativeWordStatus) IsValid() bool {
	for _, existing := range AllowedToolsCreativeWordSelectV2DataCreativeWordStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_creative_word_select_v2_data_creative_word_status value
func (v ToolsCreativeWordSelectV2DataCreativeWordStatus) Ptr() *ToolsCreativeWordSelectV2DataCreativeWordStatus {
	return &v
}
