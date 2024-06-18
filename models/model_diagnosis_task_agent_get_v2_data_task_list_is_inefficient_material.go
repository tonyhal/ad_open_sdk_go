/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial
type DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial string

// List of diagnosis_task_agent_get_v2_data_task_list_is_inefficient_material
const (
	NO_DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial      DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial = "NO"
	UNKNOWN_DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial = "UNKNOWN"
	YES_DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial     DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial = "YES"
)

// All allowed values of DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial enum
var AllowedDiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterialEnumValues = []DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial{
	"NO",
	"UNKNOWN",
	"YES",
}

// NewDiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterialFromValue returns a pointer to a valid DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterialFromValue(v string) (*DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial, error) {
	ev := DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial: valid values are %v", v, AllowedDiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_get_v2_data_task_list_is_inefficient_material value
func (v DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial) Ptr() *DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial {
	return &v
}