/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMaterialGetV30ResponseDataMaterialsInnerVideoInfo 视频信息
type AicMaterialGetV30ResponseDataMaterialsInnerVideoInfo struct {
	// 视频时长
	Duration *float64 `json:"duration,omitempty"`
	// 视频高度
	Height *int64 `json:"height,omitempty"`
	// 视频宽高比
	Ratio *float64 `json:"ratio,omitempty"`
	// 视频md5值
	Signature *string `json:"signature,omitempty"`
	// 视频ID，可以调用推送的 MAPI 推送到对应的广告账号。
	VideoId *string `json:"video_id,omitempty"`
	// 视频预览链接，可利用此链接下载视频
	VideoUrl *string `json:"video_url,omitempty"`
	// 视频宽度
	Width *int64 `json:"width,omitempty"`
}
