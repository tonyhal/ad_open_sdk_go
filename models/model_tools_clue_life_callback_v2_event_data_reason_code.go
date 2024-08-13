/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueLifeCallbackV2EventDataReasonCode
type ToolsClueLifeCallbackV2EventDataReasonCode string

// List of tools_clue_life_callback_v2_event_data_reason_code
const (
	OTHER_ToolsClueLifeCallbackV2EventDataReasonCode                        ToolsClueLifeCallbackV2EventDataReasonCode = "OTHER"
	CALL_THREE_TIMES_NO_RESPONSE_ToolsClueLifeCallbackV2EventDataReasonCode ToolsClueLifeCallbackV2EventDataReasonCode = "CALL_THREE_TIMES_NO_RESPONSE"
	INVALID_NUMBER_ToolsClueLifeCallbackV2EventDataReasonCode               ToolsClueLifeCallbackV2EventDataReasonCode = "INVALID_NUMBER"
	NO_AD_RESPONSE_ToolsClueLifeCallbackV2EventDataReasonCode               ToolsClueLifeCallbackV2EventDataReasonCode = "NO_AD_RESPONSE"
	CONNECTED_NO_INTENTION_ToolsClueLifeCallbackV2EventDataReasonCode       ToolsClueLifeCallbackV2EventDataReasonCode = "CONNECTED_NO_INTENTION"
	OFFENSIVE_LANGUAGE_ToolsClueLifeCallbackV2EventDataReasonCode           ToolsClueLifeCallbackV2EventDataReasonCode = "OFFENSIVE_LANGUAGE"
	MISMATCH_LOW_BUDGET_ToolsClueLifeCallbackV2EventDataReasonCode          ToolsClueLifeCallbackV2EventDataReasonCode = "MISMATCH_LOW_BUDGET"
	MISMATCH_DIFFERENT_LOCATIONS_ToolsClueLifeCallbackV2EventDataReasonCode ToolsClueLifeCallbackV2EventDataReasonCode = "MISMATCH_DIFFERENT_LOCATIONS"
)

// All allowed values of ToolsClueLifeCallbackV2EventDataReasonCode enum
var AllowedToolsClueLifeCallbackV2EventDataReasonCodeEnumValues = []ToolsClueLifeCallbackV2EventDataReasonCode{
	"OTHER",
	"CALL_THREE_TIMES_NO_RESPONSE",
	"INVALID_NUMBER",
	"NO_AD_RESPONSE",
	"CONNECTED_NO_INTENTION",
	"OFFENSIVE_LANGUAGE",
	"MISMATCH_LOW_BUDGET",
	"MISMATCH_DIFFERENT_LOCATIONS",
}

// NewToolsClueLifeCallbackV2EventDataReasonCodeFromValue returns a pointer to a valid ToolsClueLifeCallbackV2EventDataReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeCallbackV2EventDataReasonCodeFromValue(v string) (*ToolsClueLifeCallbackV2EventDataReasonCode, error) {
	ev := ToolsClueLifeCallbackV2EventDataReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeCallbackV2EventDataReasonCode: valid values are %v", v, AllowedToolsClueLifeCallbackV2EventDataReasonCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeCallbackV2EventDataReasonCode) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeCallbackV2EventDataReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_callback_v2_event_data_reason_code value
func (v ToolsClueLifeCallbackV2EventDataReasonCode) Ptr() *ToolsClueLifeCallbackV2EventDataReasonCode {
	return &v
}
