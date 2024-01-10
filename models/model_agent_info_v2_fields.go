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

// AgentInfoV2Fields
type AgentInfoV2Fields string

// List of agent_info_v2_fields
const (
	AGENT_ID_AgentInfoV2Fields       AgentInfoV2Fields = "agent_id"
	COMPANY_NAME_AgentInfoV2Fields   AgentInfoV2Fields = "company_name"
	CUSTOMER_ID_AgentInfoV2Fields    AgentInfoV2Fields = "customer_id"
	AGENT_NAME_AgentInfoV2Fields     AgentInfoV2Fields = "agent_name"
	COMPANY_ID_AgentInfoV2Fields     AgentInfoV2Fields = "company_id"
	CREATE_TIME_AgentInfoV2Fields    AgentInfoV2Fields = "create_time"
	CUSTOMER_NAME_AgentInfoV2Fields  AgentInfoV2Fields = "customer_name"
	ACCOUNT_STATUS_AgentInfoV2Fields AgentInfoV2Fields = "account_status"
	ROLE_AgentInfoV2Fields           AgentInfoV2Fields = "role"
)

// All allowed values of AgentInfoV2Fields enum
var AllowedAgentInfoV2FieldsEnumValues = []AgentInfoV2Fields{
	"agent_id",
	"company_name",
	"customer_id",
	"agent_name",
	"company_id",
	"create_time",
	"customer_name",
	"account_status",
	"role",
}

// NewAgentInfoV2FieldsFromValue returns a pointer to a valid AgentInfoV2Fields
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentInfoV2FieldsFromValue(v string) (*AgentInfoV2Fields, error) {
	ev := AgentInfoV2Fields(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentInfoV2Fields: valid values are %v", v, AllowedAgentInfoV2FieldsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentInfoV2Fields) IsValid() bool {
	for _, existing := range AllowedAgentInfoV2FieldsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_info_v2_fields value
func (v AgentInfoV2Fields) Ptr() *AgentInfoV2Fields {
	return &v
}
