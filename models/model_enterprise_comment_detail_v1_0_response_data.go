/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentDetailV10ResponseData
type EnterpriseCommentDetailV10ResponseData struct {
	//
	ItemCoverUrl *string `json:"item_cover_url,omitempty"`
	//
	ItemCreateTime *string `json:"item_create_time,omitempty"`
	//
	ItemDiggCount *int64 `json:"item_digg_count,omitempty"`
	//
	ItemDuration *int64 `json:"item_duration,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	ItemTitle *string `json:"item_title,omitempty"`
	//
	RepliedCommentCreateTime *string `json:"replied_comment_create_time,omitempty"`
	//
	RepliedCommentId *int64 `json:"replied_comment_id,omitempty"`
	//
	RepliedCommentImage []string `json:"replied_comment_image,omitempty"`
	//
	RepliedCommentText *string `json:"replied_comment_text,omitempty"`
	//
	RepliedOpenId *string `json:"replied_open_id,omitempty"`
}
