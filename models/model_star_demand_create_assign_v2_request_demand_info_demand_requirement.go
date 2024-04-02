/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfoDemandRequirement 产出物制作要求
type StarDemandCreateAssignV2RequestDemandInfoDemandRequirement struct {
	// 详细要求 500字内
	DetailDemand string `json:"detail_demand"`
	// 指定音乐 20字内
	Music *string `json:"music,omitempty"`
	// 其他要求 500字内
	OtherDemand *string `json:"other_demand,omitempty"`
	// 场景要求 20字内
	PropDemand *string `json:"prop_demand,omitempty"`
	// 场景要求 20字内
	SceneDemand *string `json:"scene_demand,omitempty"`
	// 文案要求 20字内
	ScriptDemand *string `json:"script_demand,omitempty"`
	// 指定话题 40字内
	SpecialTopic *string `json:"special_topic,omitempty"`
}
