/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportAdMaterialGetV10ResponseDataMaterialInfoInner struct for QianchuanReportAdMaterialGetV10ResponseDataMaterialInfoInner
type QianchuanReportAdMaterialGetV10ResponseDataMaterialInfoInner struct {
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	MaterialId *int64                                                               `json:"material_id,omitempty"`
	Metrics    *QianchuanReportAdMaterialGetV10ResponseDataMaterialInfoInnerMetrics `json:"metrics,omitempty"`
}
