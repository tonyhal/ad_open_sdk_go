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

// FilePreauditGetV30DataListMaterialChannel
type FilePreauditGetV30DataListMaterialChannel string

// List of file_preaudit_get_v3.0_data_list_material_channel
const (
	AD_FilePreauditGetV30DataListMaterialChannel       FilePreauditGetV30DataListMaterialChannel = "AD"
	DOU_PLUS_FilePreauditGetV30DataListMaterialChannel FilePreauditGetV30DataListMaterialChannel = "DOU_PLUS"
	QC_FilePreauditGetV30DataListMaterialChannel       FilePreauditGetV30DataListMaterialChannel = "QC"
)

// All allowed values of FilePreauditGetV30DataListMaterialChannel enum
var AllowedFilePreauditGetV30DataListMaterialChannelEnumValues = []FilePreauditGetV30DataListMaterialChannel{
	"AD",
	"DOU_PLUS",
	"QC",
}

// NewFilePreauditGetV30DataListMaterialChannelFromValue returns a pointer to a valid FilePreauditGetV30DataListMaterialChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditGetV30DataListMaterialChannelFromValue(v string) (*FilePreauditGetV30DataListMaterialChannel, error) {
	ev := FilePreauditGetV30DataListMaterialChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditGetV30DataListMaterialChannel: valid values are %v", v, AllowedFilePreauditGetV30DataListMaterialChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditGetV30DataListMaterialChannel) IsValid() bool {
	for _, existing := range AllowedFilePreauditGetV30DataListMaterialChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_get_v3.0_data_list_material_channel value
func (v FilePreauditGetV30DataListMaterialChannel) Ptr() *FilePreauditGetV30DataListMaterialChannel {
	return &v
}
