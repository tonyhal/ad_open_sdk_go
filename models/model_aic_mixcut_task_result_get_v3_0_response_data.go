/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMixcutTaskResultGetV30ResponseData
type AicMixcutTaskResultGetV30ResponseData struct {
	Status *AicMixcutTaskResultGetV30DataStatus `json:"status,omitempty"`
	// 视频信息
	VideoInfos []*AicMixcutTaskResultGetV30ResponseDataVideoInfosInner `json:"video_infos,omitempty"`
}
