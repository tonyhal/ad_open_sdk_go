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

// FileMaterialListV2MaterialSource
type FileMaterialListV2MaterialSource string

// List of file_material_list_v2_material_source
const (
	AD_FileMaterialListV2MaterialSource        FileMaterialListV2MaterialSource = "AD"
	QIANCHUAN_FileMaterialListV2MaterialSource FileMaterialListV2MaterialSource = "QIANCHUAN"
)

// All allowed values of FileMaterialListV2MaterialSource enum
var AllowedFileMaterialListV2MaterialSourceEnumValues = []FileMaterialListV2MaterialSource{
	"AD",
	"QIANCHUAN",
}

// NewFileMaterialListV2MaterialSourceFromValue returns a pointer to a valid FileMaterialListV2MaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialListV2MaterialSourceFromValue(v string) (*FileMaterialListV2MaterialSource, error) {
	ev := FileMaterialListV2MaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialListV2MaterialSource: valid values are %v", v, AllowedFileMaterialListV2MaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialListV2MaterialSource) IsValid() bool {
	for _, existing := range AllowedFileMaterialListV2MaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_list_v2_material_source value
func (v FileMaterialListV2MaterialSource) Ptr() *FileMaterialListV2MaterialSource {
	return &v
}
