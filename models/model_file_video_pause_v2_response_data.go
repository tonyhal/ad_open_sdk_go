/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoPauseV2ResponseData
type FileVideoPauseV2ResponseData struct {
	// 删除失败的素材id明细(素材id下仍然有至少一个的可投创意)
	FailMaterialIds []int64 `json:"fail_material_ids,omitempty"`
	// 清理结果明细
	MaterialClearResult *map[string]FileVideoPauseV2ResponseDataMaterialClearResultValue `json:"material_clear_result,omitempty"`
	// 清理结果 允许值： SUCCESS 成功 PART_SUCCESS 部分成功 FAIL 失败
	Status *string `json:"status,omitempty"`
}
