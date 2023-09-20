/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BusinessPlatformPartnerOrganizationListV2ResponseDataListInner struct for BusinessPlatformPartnerOrganizationListV2ResponseDataListInner
type BusinessPlatformPartnerOrganizationListV2ResponseDataListInner struct {
	// 发起合作组织请求的组织id
	PartnerOrganizationId *int64 `json:"partner_organization_id,omitempty"`
	// 备注，合作组织备注信息
	Remark *string                                                  `json:"remark,omitempty"`
	Status *BusinessPlatformPartnerOrganizationListV2DataListStatus `json:"status,omitempty"`
}
