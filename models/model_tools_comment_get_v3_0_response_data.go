/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30ResponseData
type ToolsCommentGetV30ResponseData struct {
	// 评论列表
	CommentList []*ToolsCommentGetV30ResponseDataCommentListInner `json:"comment_list"`
	PageInfo    ToolsCommentGetV30ResponseDataPageInfo            `json:"page_info"`
}
