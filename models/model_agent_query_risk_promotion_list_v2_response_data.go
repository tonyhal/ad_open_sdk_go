/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentQueryRiskPromotionListV2ResponseData
type AgentQueryRiskPromotionListV2ResponseData struct {
	CursorInfo *AgentQueryRiskPromotionListV2ResponseDataCursorInfo `json:"cursor_info,omitempty"`
	// 违规广告列表
	Data []*AgentQueryRiskPromotionListV2ResponseDataDataInner `json:"data,omitempty"`
}
