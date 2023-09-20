/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerVideoLocalVideo 本地视频内容
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerVideoLocalVideo struct {
	// 用户自行上传的图片url，当`local_video`不为空时，有值
	PosterUrl *string `json:"poster_url,omitempty"`
	// 视频ID，本地视频上传后得到的视频云ID，对应的是[【获取视频素材】]（https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710601820172）接口获取的视频ID，当`local_video`不为空时，有值
	VideoId *string `json:"video_id,omitempty"`
}
