/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BusinessPlatformCompanyAccountGetV30DataAccountListAccountType
type BusinessPlatformCompanyAccountGetV30DataAccountListAccountType string

// List of business_platform_company_account_get_v3.0_data_account_list_account_type
const (
	AD_BusinessPlatformCompanyAccountGetV30DataAccountListAccountType        BusinessPlatformCompanyAccountGetV30DataAccountListAccountType = "AD"
	QIANCHUAN_BusinessPlatformCompanyAccountGetV30DataAccountListAccountType BusinessPlatformCompanyAccountGetV30DataAccountListAccountType = "QIANCHUAN"
)

// All allowed values of BusinessPlatformCompanyAccountGetV30DataAccountListAccountType enum
var AllowedBusinessPlatformCompanyAccountGetV30DataAccountListAccountTypeEnumValues = []BusinessPlatformCompanyAccountGetV30DataAccountListAccountType{
	"AD",
	"QIANCHUAN",
}

// NewBusinessPlatformCompanyAccountGetV30DataAccountListAccountTypeFromValue returns a pointer to a valid BusinessPlatformCompanyAccountGetV30DataAccountListAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessPlatformCompanyAccountGetV30DataAccountListAccountTypeFromValue(v string) (*BusinessPlatformCompanyAccountGetV30DataAccountListAccountType, error) {
	ev := BusinessPlatformCompanyAccountGetV30DataAccountListAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessPlatformCompanyAccountGetV30DataAccountListAccountType: valid values are %v", v, AllowedBusinessPlatformCompanyAccountGetV30DataAccountListAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessPlatformCompanyAccountGetV30DataAccountListAccountType) IsValid() bool {
	for _, existing := range AllowedBusinessPlatformCompanyAccountGetV30DataAccountListAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to business_platform_company_account_get_v3.0_data_account_list_account_type value
func (v BusinessPlatformCompanyAccountGetV30DataAccountListAccountType) Ptr() *BusinessPlatformCompanyAccountGetV30DataAccountListAccountType {
	return &v
}
