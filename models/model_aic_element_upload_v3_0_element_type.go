/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicElementUploadV30ElementType
type AicElementUploadV30ElementType string

// List of aic_element_upload_v3.0_element_type
const (
	FOLDER_AicElementUploadV30ElementType AicElementUploadV30ElementType = "FOLDER"
	FONT_AicElementUploadV30ElementType   AicElementUploadV30ElementType = "FONT"
	IMAGE_AicElementUploadV30ElementType  AicElementUploadV30ElementType = "IMAGE"
	MUSIC_AicElementUploadV30ElementType  AicElementUploadV30ElementType = "MUSIC"
	VIDEO_AicElementUploadV30ElementType  AicElementUploadV30ElementType = "VIDEO"
	VOICE_AicElementUploadV30ElementType  AicElementUploadV30ElementType = "VOICE"
)

// Ptr returns reference to aic_element_upload_v3.0_element_type value
func (v AicElementUploadV30ElementType) Ptr() *AicElementUploadV30ElementType {
	return &v
}