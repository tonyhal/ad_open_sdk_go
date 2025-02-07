/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30DataOrdersAuditStatus
type BrandOrderListV30DataOrdersAuditStatus string

// List of brand_order_list_v3.0_data_orders_audit_status
const (
	AUDITING_BrandOrderListV30DataOrdersAuditStatus         BrandOrderListV30DataOrdersAuditStatus = "AUDITING"
	NO_CREATIVE_BrandOrderListV30DataOrdersAuditStatus      BrandOrderListV30DataOrdersAuditStatus = "NO_CREATIVE"
	PARTIALLY_PASSED_BrandOrderListV30DataOrdersAuditStatus BrandOrderListV30DataOrdersAuditStatus = "PARTIALLY_PASSED"
	PASSED_BrandOrderListV30DataOrdersAuditStatus           BrandOrderListV30DataOrdersAuditStatus = "PASSED"
	REJECTED_BrandOrderListV30DataOrdersAuditStatus         BrandOrderListV30DataOrdersAuditStatus = "REJECTED"
	UNKNOWN_BrandOrderListV30DataOrdersAuditStatus          BrandOrderListV30DataOrdersAuditStatus = "UNKNOWN"
	WAIT_AUDIT_BrandOrderListV30DataOrdersAuditStatus       BrandOrderListV30DataOrdersAuditStatus = "WAIT_AUDIT"
)

// Ptr returns reference to brand_order_list_v3.0_data_orders_audit_status value
func (v BrandOrderListV30DataOrdersAuditStatus) Ptr() *BrandOrderListV30DataOrdersAuditStatus {
	return &v
}
