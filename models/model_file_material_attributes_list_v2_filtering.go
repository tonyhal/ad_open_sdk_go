/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialAttributesListV2Filtering
type FileMaterialAttributesListV2Filtering struct {
	// 「存在搬运风险」属性最后一次更新时间，格式为yyyy-mm-dd ，筛选传入代表筛选出当天「存在搬运风险」发生过变化的素材信息
	AttributesModifyTime *string `json:"attributes_modify_time,omitempty"`
	// 结束日期，表示过滤出一段时间内上传的素材，格式为yyyy-mm-dd HH:MM:SS，如果传入起始日期，未传此参数，则默认为当前时间
	EndTime *string `json:"end_time,omitempty"`
	// 按素材ID过滤，范围为1-1000
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 素材标签过滤项，如果不传，则默认返回广告主ID下所有素材。 允许值：
	MaterialProperties []*FileMaterialAttributesListV2FilteringMaterialProperties `json:"material_properties,omitempty"`
	// 起始日期，表示按照素材上传到账户下的时间过滤，格式为yyyy-mm-dd HH:MM:SS，最早允许传入时间：2022-01-01 00:00:00 （时间是账号和素材绑定的时间，删除重绑时间会更新）
	StartTime *string `json:"start_time,omitempty"`
}
