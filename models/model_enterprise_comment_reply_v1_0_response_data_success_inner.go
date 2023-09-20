/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentReplyV10ResponseDataSuccessInner struct for EnterpriseCommentReplyV10ResponseDataSuccessInner
type EnterpriseCommentReplyV10ResponseDataSuccessInner struct {
	//
	CommentId *int64 `json:"comment_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	RepliedCommentId *int64 `json:"replied_comment_id,omitempty"`
	//
	RepliedCommentText *string `json:"replied_comment_text,omitempty"`
	//
	RepliedOpenId *string `json:"replied_open_id,omitempty"`
	//
	ReplyId *int64 `json:"reply_id,omitempty"`
	//
	Status *string `json:"status,omitempty"`
}
