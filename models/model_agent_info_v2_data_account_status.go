/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentInfoV2DataAccountStatus
type AgentInfoV2DataAccountStatus string

// List of agent_info_v2_data_account_status
const (
	STATUS_SELF_SERVICE_UNAUDITED_AgentInfoV2DataAccountStatus AgentInfoV2DataAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_CONFIRM_FAIL_AgentInfoV2DataAccountStatus           AgentInfoV2DataAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_PENDING_CONFIRM_MODIFY_AgentInfoV2DataAccountStatus AgentInfoV2DataAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_LIMIT_AgentInfoV2DataAccountStatus                  AgentInfoV2DataAccountStatus = "STATUS_LIMIT"
	STATUS_WAIT_FOR_PUBLIC_AUTH_AgentInfoV2DataAccountStatus   AgentInfoV2DataAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	STATUS_PENDING_CONFIRM_AgentInfoV2DataAccountStatus        AgentInfoV2DataAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_DISABLE_AgentInfoV2DataAccountStatus                AgentInfoV2DataAccountStatus = "STATUS_DISABLE"
	STATUS_ENABLE_AgentInfoV2DataAccountStatus                 AgentInfoV2DataAccountStatus = "STATUS_ENABLE"
	STATUS_PENDING_VERIFIED_AgentInfoV2DataAccountStatus       AgentInfoV2DataAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_CONFIRM_FAIL_END_AgentInfoV2DataAccountStatus       AgentInfoV2DataAccountStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_WAIT_FOR_BPM_AUDIT_AgentInfoV2DataAccountStatus     AgentInfoV2DataAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_CONFIRM_MODIFY_FAIL_AgentInfoV2DataAccountStatus    AgentInfoV2DataAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
)

// All allowed values of AgentInfoV2DataAccountStatus enum
var AllowedAgentInfoV2DataAccountStatusEnumValues = []AgentInfoV2DataAccountStatus{
	"STATUS_SELF_SERVICE_UNAUDITED",
	"STATUS_CONFIRM_FAIL",
	"STATUS_PENDING_CONFIRM_MODIFY",
	"STATUS_LIMIT",
	"STATUS_WAIT_FOR_PUBLIC_AUTH",
	"STATUS_PENDING_CONFIRM",
	"STATUS_DISABLE",
	"STATUS_ENABLE",
	"STATUS_PENDING_VERIFIED",
	"STATUS_CONFIRM_FAIL_END",
	"STATUS_WAIT_FOR_BPM_AUDIT",
	"STATUS_CONFIRM_MODIFY_FAIL",
}

// NewAgentInfoV2DataAccountStatusFromValue returns a pointer to a valid AgentInfoV2DataAccountStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentInfoV2DataAccountStatusFromValue(v string) (*AgentInfoV2DataAccountStatus, error) {
	ev := AgentInfoV2DataAccountStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentInfoV2DataAccountStatus: valid values are %v", v, AllowedAgentInfoV2DataAccountStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentInfoV2DataAccountStatus) IsValid() bool {
	for _, existing := range AllowedAgentInfoV2DataAccountStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_info_v2_data_account_status value
func (v AgentInfoV2DataAccountStatus) Ptr() *AgentInfoV2DataAccountStatus {
	return &v
}
