/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileVideoAwemeGetV10ResponseDataVideoListInner struct for QianchuanFileVideoAwemeGetV10ResponseDataVideoListInner
type QianchuanFileVideoAwemeGetV10ResponseDataVideoListInner struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	CommentCnt *int64 `json:"comment_cnt,omitempty"`
	//
	Duration *int64 `json:"duration,omitempty"`
	//
	Height *int32 `json:"height,omitempty"`
	//
	IsAiCreate  *bool                                                  `json:"is_ai_create,omitempty"`
	IsRecommend *QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend `json:"is_recommend,omitempty"`
	//
	LikeCnt *int64 `json:"like_cnt,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	ShareCnt *int64 `json:"share_cnt,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	ViewCnt *int64 `json:"view_cnt,omitempty"`
	//
	Width *int32 `json:"width,omitempty"`
}
