/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAudioGetV2ResponseDataListInner struct for FileAudioGetV2ResponseDataListInner
type FileAudioGetV2ResponseDataListInner struct {
	// 音频id
	AudioId *string `json:"audio_id,omitempty"`
	// 音频url
	AudioUrl *string `json:"audio_url,omitempty"`
	// 音频时长
	Duration *float64 `json:"duration,omitempty"`
	// 音频文件名称
	FileName *string `json:"file_name,omitempty"`
	// 音频素材id
	MaterialId *int64 `json:"material_id,omitempty"`
	// 音频md5值
	Signature *string                       `json:"signature,omitempty"`
	Source    *FileAudioGetV2DataListSource `json:"source,omitempty"`
}
