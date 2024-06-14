/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial
type DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial string

// List of diagnosis_task_agent_get_v2_data_task_list_is_ecp_high_quality_material
const (
	NO_DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial      DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial = "NO"
	UNKNOWN_DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial = "UNKNOWN"
	YES_DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial     DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial = "YES"
)

// All allowed values of DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial enum
var AllowedDiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterialEnumValues = []DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial{
	"NO",
	"UNKNOWN",
	"YES",
}

// NewDiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterialFromValue returns a pointer to a valid DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterialFromValue(v string) (*DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial, error) {
	ev := DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial: valid values are %v", v, AllowedDiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_get_v2_data_task_list_is_ecp_high_quality_material value
func (v DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial) Ptr() *DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial {
	return &v
}
