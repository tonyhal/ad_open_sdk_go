/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10ResponseData
type EnterpriseCommentListGetV10ResponseData struct {
	//
	CommentList []*EnterpriseCommentListGetV10ResponseDataCommentListInner `json:"comment_list,omitempty"`
	PageInfo    *EnterpriseCommentListGetV10ResponseDataPageInfo           `json:"page_info,omitempty"`
}
