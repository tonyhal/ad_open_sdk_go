/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPreAuditSendV2Request struct for ToolsPreAuditSendV2Request
type ToolsPreAuditSendV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 前置预审物料列表
	PreAuditMaterials []*ToolsPreAuditSendV2RequestPreAuditMaterialsInner `json:"pre_audit_materials"`
}
