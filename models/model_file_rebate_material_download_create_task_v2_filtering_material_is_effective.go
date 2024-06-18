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

// FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective
type FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective string

// List of file_rebate_material_download_create_task_v2_filtering_material_is_effective
const (
	NO_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective  FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective = "NO"
	YES_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective = "YES"
)

// All allowed values of FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective enum
var AllowedFileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffectiveEnumValues = []FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective{
	"NO",
	"YES",
}

// NewFileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffectiveFromValue returns a pointer to a valid FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffectiveFromValue(v string) (*FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective, error) {
	ev := FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective: valid values are %v", v, AllowedFileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffectiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective) IsValid() bool {
	for _, existing := range AllowedFileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffectiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_material_is_effective value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective {
	return &v
}