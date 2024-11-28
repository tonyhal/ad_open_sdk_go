/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskDetailV30ResponseDataStarTaskAlbumMicroAppInfo 星广联投短剧版任务小程序相关信息
type StardeliveryTaskDetailV30ResponseDataStarTaskAlbumMicroAppInfo struct {
	// 星图任务内的字节小程序ID
	MicroAppId *string `json:"micro_app_id,omitempty"`
	// 字节小程序名称，即星图任务内的组件标题
	MicroAppName *string `json:"micro_app_name,omitempty"`
	// 小程序页面地址
	MicroAppPath *string `json:"micro_app_path,omitempty"`
	// 短剧剧目id
	StarTaskAlbumId *string `json:"star_task_album_id,omitempty"`
	// 短剧剧目名称
	StarTaskAlbumName *string `json:"star_task_album_name,omitempty"`
}