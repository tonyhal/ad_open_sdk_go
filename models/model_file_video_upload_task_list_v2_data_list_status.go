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

// FileVideoUploadTaskListV2DataListStatus
type FileVideoUploadTaskListV2DataListStatus string

// List of file_video_upload_task_list_v2_data_list_status
const (
	PROCESS_FileVideoUploadTaskListV2DataListStatus FileVideoUploadTaskListV2DataListStatus = "PROCESS"
	SUCCESS_FileVideoUploadTaskListV2DataListStatus FileVideoUploadTaskListV2DataListStatus = "SUCCESS"
	FAILED_FileVideoUploadTaskListV2DataListStatus  FileVideoUploadTaskListV2DataListStatus = "FAILED"
)

// All allowed values of FileVideoUploadTaskListV2DataListStatus enum
var AllowedFileVideoUploadTaskListV2DataListStatusEnumValues = []FileVideoUploadTaskListV2DataListStatus{
	"PROCESS",
	"SUCCESS",
	"FAILED",
}

// NewFileVideoUploadTaskListV2DataListStatusFromValue returns a pointer to a valid FileVideoUploadTaskListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoUploadTaskListV2DataListStatusFromValue(v string) (*FileVideoUploadTaskListV2DataListStatus, error) {
	ev := FileVideoUploadTaskListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoUploadTaskListV2DataListStatus: valid values are %v", v, AllowedFileVideoUploadTaskListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoUploadTaskListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedFileVideoUploadTaskListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_upload_task_list_v2_data_list_status value
func (v FileVideoUploadTaskListV2DataListStatus) Ptr() *FileVideoUploadTaskListV2DataListStatus {
	return &v
}
