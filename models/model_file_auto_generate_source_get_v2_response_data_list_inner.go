/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAutoGenerateSourceGetV2ResponseDataListInner struct for FileAutoGenerateSourceGetV2ResponseDataListInner
type FileAutoGenerateSourceGetV2ResponseDataListInner struct {
	//
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	SourceMaterialIds []int64 `json:"source_material_ids,omitempty"`
}
