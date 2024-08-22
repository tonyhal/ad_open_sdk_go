/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag
type FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag string

// List of file_rebate_material_download_create_task_v2_filtering_material_tag
const (
	HIGH_QUALITY_MATERIAL_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag    FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag = "HIGH_QUALITY_MATERIAL"
	LOW_QUALITY_MATERIAL_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag     FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag = "LOW_QUALITY_MATERIAL"
	FIRST_EFFECTIVE_MATERIAL_FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag = "FIRST_EFFECTIVE_MATERIAL"
)

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_material_tag value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag {
	return &v
}
