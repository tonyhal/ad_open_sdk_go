/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestTrackUrlInfo 第三方监测链接
type BrandCreativeCreateV30RequestTrackUrlInfo struct {
	// 点击监测
	ActionTrackUrlList []string `json:"action_track_url_list,omitempty"`
	// 视频5s播放监测
	EffectiveFrameUrlList []string `json:"effective_frame_url_list,omitempty"`
	// 展示监测
	TrackUrlList     []string                                           `json:"track_url_list,omitempty"`
	TrackUrlSendType BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType `json:"track_url_send_type"`
}
