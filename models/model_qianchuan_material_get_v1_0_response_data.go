/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanMaterialGetV10ResponseData
type QianchuanMaterialGetV10ResponseData struct {
	// 返回的素材信息列表
	AdMaterialInfos []*QianchuanMaterialGetV10ResponseDataAdMaterialInfosInner `json:"ad_material_infos"`
	PageInfo        *QianchuanMaterialGetV10ResponseDataPageInfo               `json:"page_info,omitempty"`
}