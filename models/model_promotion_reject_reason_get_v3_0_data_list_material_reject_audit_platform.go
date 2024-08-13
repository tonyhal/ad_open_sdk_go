/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform
type PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform string

// List of promotion_reject_reason_get_v3.0_data_list_material_reject_audit_platform
const (
	AD_PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform      PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform = "AD"
	CONTENT_PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform = "CONTENT"
)

// All allowed values of PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform enum
var AllowedPromotionRejectReasonGetV30DataListMaterialRejectAuditPlatformEnumValues = []PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform{
	"AD",
	"CONTENT",
}

// NewPromotionRejectReasonGetV30DataListMaterialRejectAuditPlatformFromValue returns a pointer to a valid PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionRejectReasonGetV30DataListMaterialRejectAuditPlatformFromValue(v string) (*PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform, error) {
	ev := PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform: valid values are %v", v, AllowedPromotionRejectReasonGetV30DataListMaterialRejectAuditPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform) IsValid() bool {
	for _, existing := range AllowedPromotionRejectReasonGetV30DataListMaterialRejectAuditPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_reject_reason_get_v3.0_data_list_material_reject_audit_platform value
func (v PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform) Ptr() *PromotionRejectReasonGetV30DataListMaterialRejectAuditPlatform {
	return &v
}
