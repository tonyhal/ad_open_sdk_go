/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostItemGroupDetailV2ResponseData
type StarVasGetBoostItemGroupDetailV2ResponseData struct {
	AllFlowStatInfo   StarVasGetBoostItemGroupDetailV2ResponseDataAllFlowStatInfo   `json:"all_flow_stat_info"`
	BoostFlowStatInfo StarVasGetBoostItemGroupDetailV2ResponseDataBoostFlowStatInfo `json:"boost_flow_stat_info"`
	// 指派单业务指标，最多返回50个指派单的信息
	OrderStatInfo []*StarVasGetBoostItemGroupDetailV2ResponseDataOrderStatInfoInner `json:"order_stat_info"`
	TaskInfo      StarVasGetBoostItemGroupDetailV2ResponseDataTaskInfo              `json:"task_info"`
}
