/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostGroupListV2ResponseData
type StarVasGetBoostGroupListV2ResponseData struct {
	// 助推组列表
	BoostGroupInfos []*StarVasGetBoostGroupListV2ResponseDataBoostGroupInfosInner `json:"boost_group_infos"`
	Pagination      StarVasGetBoostGroupListV2ResponseDataPagination              `json:"pagination"`
}
