/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexRemarkV2ResponseDataListInner struct for ToolsRubeexRemarkV2ResponseDataListInner
type ToolsRubeexRemarkV2ResponseDataListInner struct {
	// data_md5指rubeex工具产出的试玩素材作品版本，一个rubeex试玩素材可能有多个作品版本
	DataMd5 *string `json:"data_md5,omitempty"`
	// 作品id即试玩素材id，属于试玩素材唯一标识
	ProjectId *float64 `json:"project_id,omitempty"`
	// 场景
	Section *string `json:"section,omitempty"`
	// 区域
	SectionArea *string `json:"section_area,omitempty"`
	// 区域名称
	SectionAreaRemark *string `json:"section_area_remark,omitempty"`
	// 场景名称
	SectionRemark *string `json:"section_remark,omitempty"`
}