/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BusinessPlatformCompanyInfoGetV30ResponseDataCompanyInfoInner struct for BusinessPlatformCompanyInfoGetV30ResponseDataCompanyInfoInner
type BusinessPlatformCompanyInfoGetV30ResponseDataCompanyInfoInner struct {
	//
	CompanyId *int64 `json:"company_id,omitempty"`
	//
	CompanyName *string                                                 `json:"company_name,omitempty"`
	Status      *BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus `json:"status,omitempty"`
	Type        *BusinessPlatformCompanyInfoGetV30DataCompanyInfoType   `json:"type,omitempty"`
}
