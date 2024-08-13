/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAgentV2ResponseDataVideoInfo 视频信息
type FileVideoAgentV2ResponseDataVideoInfo struct {
	// 视频播放时长
	Duration *float64 `json:"duration,omitempty"`
	// 视频宽度
	Height *int64 `json:"height,omitempty"`
	// 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	MaterialId *int64 `json:"material_id,omitempty"`
	// 视频大小
	Size *int64 `json:"size,omitempty"`
	// 视频唯一标识vid
	VideoId *string `json:"video_id,omitempty"`
	// 视频md5
	VideoSignature *string `json:"video_signature,omitempty"`
	// 视频播放url
	VideoUrl *string `json:"video_url,omitempty"`
	// 视频高度
	Width *int64 `json:"width,omitempty"`
}
