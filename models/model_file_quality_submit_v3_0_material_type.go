/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileQualitySubmitV30MaterialType
type FileQualitySubmitV30MaterialType string

// List of file_quality_submit_v3.0_material_type
const (
	VIDEO_FileQualitySubmitV30MaterialType FileQualitySubmitV30MaterialType = "VIDEO"
)

// All allowed values of FileQualitySubmitV30MaterialType enum
var AllowedFileQualitySubmitV30MaterialTypeEnumValues = []FileQualitySubmitV30MaterialType{
	"VIDEO",
}

// NewFileQualitySubmitV30MaterialTypeFromValue returns a pointer to a valid FileQualitySubmitV30MaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileQualitySubmitV30MaterialTypeFromValue(v string) (*FileQualitySubmitV30MaterialType, error) {
	ev := FileQualitySubmitV30MaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileQualitySubmitV30MaterialType: valid values are %v", v, AllowedFileQualitySubmitV30MaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileQualitySubmitV30MaterialType) IsValid() bool {
	for _, existing := range AllowedFileQualitySubmitV30MaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_quality_submit_v3.0_material_type value
func (v FileQualitySubmitV30MaterialType) Ptr() *FileQualitySubmitV30MaterialType {
	return &v
}