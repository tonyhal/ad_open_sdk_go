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

// FilePreauditSubmitV30MaterialType
type FilePreauditSubmitV30MaterialType string

// List of file_preaudit_submit_v3.0_material_type
const (
	VIDEO_FilePreauditSubmitV30MaterialType FilePreauditSubmitV30MaterialType = "VIDEO"
)

// All allowed values of FilePreauditSubmitV30MaterialType enum
var AllowedFilePreauditSubmitV30MaterialTypeEnumValues = []FilePreauditSubmitV30MaterialType{
	"VIDEO",
}

// NewFilePreauditSubmitV30MaterialTypeFromValue returns a pointer to a valid FilePreauditSubmitV30MaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditSubmitV30MaterialTypeFromValue(v string) (*FilePreauditSubmitV30MaterialType, error) {
	ev := FilePreauditSubmitV30MaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditSubmitV30MaterialType: valid values are %v", v, AllowedFilePreauditSubmitV30MaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditSubmitV30MaterialType) IsValid() bool {
	for _, existing := range AllowedFilePreauditSubmitV30MaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_submit_v3.0_material_type value
func (v FilePreauditSubmitV30MaterialType) Ptr() *FilePreauditSubmitV30MaterialType {
	return &v
}
