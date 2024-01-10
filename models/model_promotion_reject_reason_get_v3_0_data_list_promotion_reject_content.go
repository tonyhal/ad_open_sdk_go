/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionRejectReasonGetV30DataListPromotionRejectContent
type PromotionRejectReasonGetV30DataListPromotionRejectContent string

// List of promotion_reject_reason_get_v3.0_data_list_promotion_reject_content
const (
	Enum_下载链接_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "下载链接"
	Enum_副标题_PromotionRejectReasonGetV30DataListPromotionRejectContent      PromotionRejectReasonGetV30DataListPromotionRejectContent = "副标题"
	Enum_卡片主图_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "卡片主图"
	Enum_卡片标题_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "卡片标题"
	Enum_头条group_PromotionRejectReasonGetV30DataListPromotionRejectContent  PromotionRejectReasonGetV30DataListPromotionRejectContent = "头条group"
	Enum_小程序_PromotionRejectReasonGetV30DataListPromotionRejectContent      PromotionRejectReasonGetV30DataListPromotionRejectContent = "小程序"
	Enum_应用包名_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "应用包名"
	Enum_应用名_PromotionRejectReasonGetV30DataListPromotionRejectContent      PromotionRejectReasonGetV30DataListPromotionRejectContent = "应用名"
	Enum_应用详情页_PromotionRejectReasonGetV30DataListPromotionRejectContent    PromotionRejectReasonGetV30DataListPromotionRejectContent = "应用详情页"
	Enum_录入文章_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "录入文章"
	Enum_抖音头像_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "抖音头像"
	Enum_抖音昵称_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "抖音昵称"
	Enum_按钮附加创意_PromotionRejectReasonGetV30DataListPromotionRejectContent   PromotionRejectReasonGetV30DataListPromotionRejectContent = "按钮附加创意"
	Enum_推广卖点_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "推广卖点"
	Enum_文章链接_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "文章链接"
	Enum_来源_PromotionRejectReasonGetV30DataListPromotionRejectContent       PromotionRejectReasonGetV30DataListPromotionRejectContent = "来源"
	Enum_标题_PromotionRejectReasonGetV30DataListPromotionRejectContent       PromotionRejectReasonGetV30DataListPromotionRejectContent = "标题"
	Enum_电话号码附加创意_PromotionRejectReasonGetV30DataListPromotionRejectContent PromotionRejectReasonGetV30DataListPromotionRejectContent = "电话号码附加创意"
	Enum_直达链接_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "直达链接"
	Enum_落地页_PromotionRejectReasonGetV30DataListPromotionRejectContent      PromotionRejectReasonGetV30DataListPromotionRejectContent = "落地页"
	Enum_表单附加创意_PromotionRejectReasonGetV30DataListPromotionRejectContent   PromotionRejectReasonGetV30DataListPromotionRejectContent = "表单附加创意"
	Enum_视频锚点_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "视频锚点"
	Enum_计划_PromotionRejectReasonGetV30DataListPromotionRejectContent       PromotionRejectReasonGetV30DataListPromotionRejectContent = "计划"
	Enum_附加创意_PromotionRejectReasonGetV30DataListPromotionRejectContent     PromotionRejectReasonGetV30DataListPromotionRejectContent = "附加创意"
)

// All allowed values of PromotionRejectReasonGetV30DataListPromotionRejectContent enum
var AllowedPromotionRejectReasonGetV30DataListPromotionRejectContentEnumValues = []PromotionRejectReasonGetV30DataListPromotionRejectContent{
	"下载链接",
	"副标题",
	"卡片主图",
	"卡片标题",
	"头条group",
	"小程序",
	"应用包名",
	"应用名",
	"应用详情页",
	"录入文章",
	"抖音头像",
	"抖音昵称",
	"按钮附加创意",
	"推广卖点",
	"文章链接",
	"来源",
	"标题",
	"电话号码附加创意",
	"直达链接",
	"落地页",
	"表单附加创意",
	"视频锚点",
	"计划",
	"附加创意",
}

// NewPromotionRejectReasonGetV30DataListPromotionRejectContentFromValue returns a pointer to a valid PromotionRejectReasonGetV30DataListPromotionRejectContent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionRejectReasonGetV30DataListPromotionRejectContentFromValue(v string) (*PromotionRejectReasonGetV30DataListPromotionRejectContent, error) {
	ev := PromotionRejectReasonGetV30DataListPromotionRejectContent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionRejectReasonGetV30DataListPromotionRejectContent: valid values are %v", v, AllowedPromotionRejectReasonGetV30DataListPromotionRejectContentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionRejectReasonGetV30DataListPromotionRejectContent) IsValid() bool {
	for _, existing := range AllowedPromotionRejectReasonGetV30DataListPromotionRejectContentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_reject_reason_get_v3.0_data_list_promotion_reject_content value
func (v PromotionRejectReasonGetV30DataListPromotionRejectContent) Ptr() *PromotionRejectReasonGetV30DataListPromotionRejectContent {
	return &v
}
