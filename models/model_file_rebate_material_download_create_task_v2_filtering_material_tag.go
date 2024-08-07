/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag
type FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag string

// List of file_rebate_material_download_create_task_v2_filtering_material_tag
const (
	HIGH_QUALITY_MATERIAL_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag    FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag = "HIGH_QUALITY_MATERIAL"
	LOW_QUALITY_MATERIAL_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag     FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag = "LOW_QUALITY_MATERIAL"
	FIRST_EFFECTIVE_MATERIAL_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag = "FIRST_EFFECTIVE_MATERIAL"
)

// All allowed values of FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag enum
var AllowedFileRebateMaterialDownloadCreateTaskV2FilteringMaterialTagEnumValues = []FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag{
	"HIGH_QUALITY_MATERIAL",
	"LOW_QUALITY_MATERIAL",
	"FIRST_EFFECTIVE_MATERIAL",
}

// NewFileRebateMaterialDownloadCreateTaskV2FilteringMaterialTagFromValue returns a pointer to a valid FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileRebateMaterialDownloadCreateTaskV2FilteringMaterialTagFromValue(v string) (*FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag, error) {
	ev := FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag: valid values are %v", v, AllowedFileRebateMaterialDownloadCreateTaskV2FilteringMaterialTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag) IsValid() bool {
	for _, existing := range AllowedFileRebateMaterialDownloadCreateTaskV2FilteringMaterialTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_material_tag value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag {
	return &v
}
