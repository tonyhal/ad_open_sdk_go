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

// FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType
type FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType string

// List of file_rebate_material_download_create_task_v2_filtering_rebate_calc_policy_type
const (
	NORMAL_POLICY_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType    FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType = "NORMAL_POLICY"
	EXCLUSIVE_POLICY_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType = "EXCLUSIVE_POLICY"
	CAR_POLICY_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType       FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType = "CAR_POLICY"
)

// All allowed values of FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType enum
var AllowedFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyTypeEnumValues = []FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType{
	"NORMAL_POLICY",
	"EXCLUSIVE_POLICY",
	"CAR_POLICY",
}

// NewFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyTypeFromValue returns a pointer to a valid FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyTypeFromValue(v string) (*FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType, error) {
	ev := FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType: valid values are %v", v, AllowedFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType) IsValid() bool {
	for _, existing := range AllowedFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_rebate_calc_policy_type value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType {
	return &v
}
