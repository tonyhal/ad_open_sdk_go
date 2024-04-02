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

// FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes
type FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes string

// List of file_video_material_clear_task_get_v2_data_list_clear_task_params_clear_material_types
const (
	INEFFICIENT_MATERIAL_FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes   FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes       FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes = "SIMILAR_QUEUE_MATERIAL"
)

// All allowed values of FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes enum
var AllowedFileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypesEnumValues = []FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes{
	"INEFFICIENT_MATERIAL",
	"SIMILAR_MATERIAL",
	"SIMILAR_QUEUE_MATERIAL",
}

// NewFileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypesFromValue returns a pointer to a valid FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypesFromValue(v string) (*FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes, error) {
	ev := FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes: valid values are %v", v, AllowedFileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_get_v2_data_list_clear_task_params_clear_material_types value
func (v FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes) Ptr() *FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes {
	return &v
}
