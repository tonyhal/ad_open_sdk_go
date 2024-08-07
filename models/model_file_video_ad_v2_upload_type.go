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

// FileVideoAdV2UploadType
type FileVideoAdV2UploadType string

// List of file_video_ad_v2_upload_type
const (
	UPLOAD_BY_FILE_FileVideoAdV2UploadType FileVideoAdV2UploadType = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL_FileVideoAdV2UploadType  FileVideoAdV2UploadType = "UPLOAD_BY_URL"
)

// All allowed values of FileVideoAdV2UploadType enum
var AllowedFileVideoAdV2UploadTypeEnumValues = []FileVideoAdV2UploadType{
	"UPLOAD_BY_FILE",
	"UPLOAD_BY_URL",
}

// NewFileVideoAdV2UploadTypeFromValue returns a pointer to a valid FileVideoAdV2UploadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoAdV2UploadTypeFromValue(v string) (*FileVideoAdV2UploadType, error) {
	ev := FileVideoAdV2UploadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoAdV2UploadType: valid values are %v", v, AllowedFileVideoAdV2UploadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoAdV2UploadType) IsValid() bool {
	for _, existing := range AllowedFileVideoAdV2UploadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_ad_v2_upload_type value
func (v FileVideoAdV2UploadType) Ptr() *FileVideoAdV2UploadType {
	return &v
}
