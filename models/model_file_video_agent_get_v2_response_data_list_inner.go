/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAgentGetV2ResponseDataListInner struct for FileVideoAgentGetV2ResponseDataListInner
type FileVideoAgentGetV2ResponseDataListInner struct {
	// 素材的上传时间，格式：\"yyyy-mm-dd HH:MM:SS\"
	CreateTime *string `json:"create_time,omitempty"`
	// 素材的文件名
	Filename *string `json:"filename,omitempty"`
	// 视频ID
	Id *string `json:"id,omitempty"`
	// 视频md5值
	Signature *string `json:"signature,omitempty"`
	// 素材来源，详见【附录-素材来源】
	Source *string `json:"source,omitempty"`
	// 视频地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL” 链接仅做预览使用，预览链接有效期为1小时
	Url *string `json:"url,omitempty"`
}
