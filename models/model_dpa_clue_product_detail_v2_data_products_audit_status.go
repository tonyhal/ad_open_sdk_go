/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductDetailV2DataProductsAuditStatus
type DpaClueProductDetailV2DataProductsAuditStatus string

// List of dpa_clue_product_detail_v2_data_products_audit_status
const (
	AUDIT_STATUS_APPROVE_DpaClueProductDetailV2DataProductsAuditStatus DpaClueProductDetailV2DataProductsAuditStatus = "AUDIT_STATUS_APPROVE"
	AUDIT_STATUS_INIT_DpaClueProductDetailV2DataProductsAuditStatus    DpaClueProductDetailV2DataProductsAuditStatus = "AUDIT_STATUS_INIT"
	AUDIT_STATUS_REJECT_DpaClueProductDetailV2DataProductsAuditStatus  DpaClueProductDetailV2DataProductsAuditStatus = "AUDIT_STATUS_REJECT"
)

// Ptr returns reference to dpa_clue_product_detail_v2_data_products_audit_status value
func (v DpaClueProductDetailV2DataProductsAuditStatus) Ptr() *DpaClueProductDetailV2DataProductsAuditStatus {
	return &v
}
