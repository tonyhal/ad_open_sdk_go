/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetInfoV2ResponseDataOrderListInnerAuthorInfo 达人信息
type StarOrderGetInfoV2ResponseDataOrderListInnerAuthorInfo struct {
	// 达人ID
	AuthorId *int64 `json:"author_id,omitempty"`
	// 达人名称
	AuthorName *string `json:"author_name,omitempty"`
	// 头像地址
	AvatarUri *string `json:"avatar_uri,omitempty"`
	// 联系电话
	ContactPhone *string `json:"contact_phone,omitempty"`
}
