/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentReplyGetV30ResponseData
type ToolsCommentReplyGetV30ResponseData struct {
	PageInfo ToolsCommentReplyGetV30ResponseDataPageInfo `json:"page_info"`
	// 回复列表
	ReplyList []*ToolsCommentReplyGetV30ResponseDataReplyListInner `json:"reply_list"`
}
