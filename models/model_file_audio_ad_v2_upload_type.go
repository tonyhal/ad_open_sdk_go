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

// FileAudioAdV2UploadType
type FileAudioAdV2UploadType string

// List of file_audio_ad_v2_upload_type
const (
	UPLOAD_BY_FILE_FileAudioAdV2UploadType FileAudioAdV2UploadType = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL_FileAudioAdV2UploadType  FileAudioAdV2UploadType = "UPLOAD_BY_URL"
)

// All allowed values of FileAudioAdV2UploadType enum
var AllowedFileAudioAdV2UploadTypeEnumValues = []FileAudioAdV2UploadType{
	"UPLOAD_BY_FILE",
	"UPLOAD_BY_URL",
}

// NewFileAudioAdV2UploadTypeFromValue returns a pointer to a valid FileAudioAdV2UploadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileAudioAdV2UploadTypeFromValue(v string) (*FileAudioAdV2UploadType, error) {
	ev := FileAudioAdV2UploadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileAudioAdV2UploadType: valid values are %v", v, AllowedFileAudioAdV2UploadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileAudioAdV2UploadType) IsValid() bool {
	for _, existing := range AllowedFileAudioAdV2UploadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_audio_ad_v2_upload_type value
func (v FileAudioAdV2UploadType) Ptr() *FileAudioAdV2UploadType {
	return &v
}
