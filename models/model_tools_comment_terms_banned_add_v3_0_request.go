/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentTermsBannedAddV30Request struct for ToolsCommentTermsBannedAddV30Request
type ToolsCommentTermsBannedAddV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 抖音号，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeId *string `json:"aweme_id,omitempty"`
	// 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv *bool `json:"is_apply_to_adv,omitempty"`
	// 屏蔽词列表，若添加的屏蔽词已存在，不会再次新增，同一个屏蔽词只会在屏蔽词中记录一次 一次最多添加100个，单个屏蔽词长度范围为0-20字
	Terms []string `json:"terms"`
}