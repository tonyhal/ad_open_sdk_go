/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostItemGroupDetailV2ResponseDataOrderStatInfoInnerAllFlowStatInfo 全流量业务指标
type StarVasGetBoostItemGroupDetailV2ResponseDataOrderStatInfoInnerAllFlowStatInfo struct {
	// 评论数
	Comment int64 `json:"comment"`
	// 组件点击数，仅BoostAction=LINK_ACTION有效
	ComponentClickPv *int64 `json:"component_click_pv,omitempty"`
	// 组件CTR，百分数*100取整，如:1% -> 100，仅BoostAction=LINK_ACTION有效
	ComponentCtr *int64 `json:"component_ctr,omitempty"`
	// 转化数，仅BoostAction=LINK_ACTION有效
	ConvertCount *int64 `json:"convert_count,omitempty"`
	// 业务目标人群实际覆盖数，仅BoostType=CUSTOM_PACKAGE&& BoostAction=NATIVE_ACTION有效
	FinishedWatcherUv *int64 `json:"finished_watcher_uv,omitempty"`
	// 点赞数
	Like int64 `json:"like"`
	// 播放数
	Play int64 `json:"play"`
	// 五秒播放率
	PlayRateWith5s int64 `json:"play_rate_with5s"`
	// 分享数
	Share int64 `json:"share"`
}
