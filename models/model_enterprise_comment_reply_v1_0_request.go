/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentReplyV10Request struct for EnterpriseCommentReplyV10Request
type EnterpriseCommentReplyV10Request struct {
	//
	CcAccountId int64 `json:"cc_account_id"`
	//
	CommentIds []int64 `json:"comment_ids"`
	//
	EDouyinId string `json:"e_douyin_id"`
	//
	ReplyText string `json:"reply_text"`
}
