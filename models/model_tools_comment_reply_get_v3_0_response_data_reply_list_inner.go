/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentReplyGetV30ResponseDataReplyListInner struct for ToolsCommentReplyGetV30ResponseDataReplyListInner
type ToolsCommentReplyGetV30ResponseDataReplyListInner struct {
	// 评论用户抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 评论用户昵称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 评论id
	CommentId *int64 `json:"comment_id,omitempty"`
	// 创建时间
	CreateTime string                                          `json:"create_time"`
	HideStatus *ToolsCommentReplyGetV30DataReplyListHideStatus `json:"hide_status,omitempty"`
	// 该评论是否置顶，0：表示不置顶，1：表示置顶
	IsStick *int64 `json:"is_stick,omitempty"`
	// 抖音视频id
	ItemId *int64 `json:"item_id,omitempty"`
	// 点赞数
	LikeCount *int64 `json:"like_count,omitempty"`
	// 评论内容
	Text *string `json:"text,omitempty"`
}
