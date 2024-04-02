/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmUpdateChallengeV2RequestChallengeInfo
type StarDemandOmUpdateChallengeV2RequestChallengeInfo struct {
	// 组件标题
	AnchorTitle *string `json:"anchor_title,omitempty"`
	// 达人侧任务名称
	AuthorTaskName *string `json:"author_task_name,omitempty"`
	// 任务介绍 文本，最长500字。如果不传递参数，那么字段将保留原始任务信息，不会进行更新
	DemandDesc *string `json:"demand_desc,omitempty"`
	// 任务标签
	OmTaskTag []string `json:"om_task_tag,omitempty"`
	// 示例视频id 最多5个
	SampleVideo []int64 `json:"sample_video,omitempty"`
	// 小程序落地页地址 字符串，支持校验合规性
	StartPage *string `json:"start_page,omitempty"`
	// 任务头图 URL文本
	TaskHeadImage *string `json:"task_head_image,omitempty"`
	// 任务图标 URL文本
	TaskIcon *string `json:"task_icon,omitempty"`
}
