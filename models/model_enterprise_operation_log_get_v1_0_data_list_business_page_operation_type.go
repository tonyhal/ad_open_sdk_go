/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseOperationLogGetV10DataListBusinessPageOperationType
type EnterpriseOperationLogGetV10DataListBusinessPageOperationType string

// List of enterprise_operation_log_get_v1.0_data_list_business_page_operation_type
const (
	AD_EnterpriseOperationLogGetV10DataListBusinessPageOperationType    EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "AD"
	VIDEO_EnterpriseOperationLogGetV10DataListBusinessPageOperationType EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "VIDEO"
	LIVE_EnterpriseOperationLogGetV10DataListBusinessPageOperationType  EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "LIVE"
	DOU_EnterpriseOperationLogGetV10DataListBusinessPageOperationType   EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "DOU"
)

// Ptr returns reference to enterprise_operation_log_get_v1.0_data_list_business_page_operation_type value
func (v EnterpriseOperationLogGetV10DataListBusinessPageOperationType) Ptr() *EnterpriseOperationLogGetV10DataListBusinessPageOperationType {
	return &v
}
