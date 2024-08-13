/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentReplyV30Request struct for ToolsCommentReplyV30Request
type ToolsCommentReplyV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 评论id列表，只允许传入小于等于20个评论id
	CommentIds []int64 `json:"comment_ids"`
	// 回复内容，最多支持100字
	ReplyText string `json:"reply_text"`
}
