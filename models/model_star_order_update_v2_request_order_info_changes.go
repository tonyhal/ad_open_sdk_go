/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderUpdateV2RequestOrderInfoChanges 任务信息变动
type StarOrderUpdateV2RequestOrderInfoChanges struct {
	// 组件点击监测链接
	ComponentClickMonitorUrl *string `json:"component_click_monitor_url,omitempty"`
	// 视频曝光监测链接
	ItemShowMonitorUrl *string `json:"item_show_monitor_url,omitempty"`
	// 常规组件ID列表（最多一个）
	LinkComponentIds []int64 `json:"link_component_ids,omitempty"`
}
