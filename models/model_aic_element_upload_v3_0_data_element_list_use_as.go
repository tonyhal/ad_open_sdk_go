/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicElementUploadV30DataElementListUseAs
type AicElementUploadV30DataElementListUseAs string

// List of aic_element_upload_v3.0_data_element_list_use_as
const (
	BACKGROUND_AicElementUploadV30DataElementListUseAs AicElementUploadV30DataElementListUseAs = "BACKGROUND"
	FOREGROUND_AicElementUploadV30DataElementListUseAs AicElementUploadV30DataElementListUseAs = "FOREGROUND"
	POST_VIDEO_AicElementUploadV30DataElementListUseAs AicElementUploadV30DataElementListUseAs = "POST_VIDEO"
	PRE_VIDEO_AicElementUploadV30DataElementListUseAs  AicElementUploadV30DataElementListUseAs = "PRE_VIDEO"
	RAW_AicElementUploadV30DataElementListUseAs        AicElementUploadV30DataElementListUseAs = "RAW"
)

// Ptr returns reference to aic_element_upload_v3.0_data_element_list_use_as value
func (v AicElementUploadV30DataElementListUseAs) Ptr() *AicElementUploadV30DataElementListUseAs {
	return &v
}
