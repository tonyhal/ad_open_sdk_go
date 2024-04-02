/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30ResponseData
type AdvertiserDeliveryPkgSubmitV30ResponseData struct {
	// 推广产品组id，可用于后续的查询或编辑
	PkgId *int64 `json:"pkg_id,omitempty"`
	// 系统生成的资质id，每份资质对应一个id
	QualificationIds []int64 `json:"qualification_ids,omitempty"`
}
