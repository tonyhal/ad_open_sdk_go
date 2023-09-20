/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BusinessPlatformCompanyAccountGetV30AccountType
type BusinessPlatformCompanyAccountGetV30AccountType string

// List of business_platform_company_account_get_v3.0_account_type
const (
	AD_BusinessPlatformCompanyAccountGetV30AccountType        BusinessPlatformCompanyAccountGetV30AccountType = "AD"
	QIANCHUAN_BusinessPlatformCompanyAccountGetV30AccountType BusinessPlatformCompanyAccountGetV30AccountType = "QIANCHUAN"
)

// All allowed values of BusinessPlatformCompanyAccountGetV30AccountType enum
var AllowedBusinessPlatformCompanyAccountGetV30AccountTypeEnumValues = []BusinessPlatformCompanyAccountGetV30AccountType{
	"AD",
	"QIANCHUAN",
}

// NewBusinessPlatformCompanyAccountGetV30AccountTypeFromValue returns a pointer to a valid BusinessPlatformCompanyAccountGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessPlatformCompanyAccountGetV30AccountTypeFromValue(v string) (*BusinessPlatformCompanyAccountGetV30AccountType, error) {
	ev := BusinessPlatformCompanyAccountGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessPlatformCompanyAccountGetV30AccountType: valid values are %v", v, AllowedBusinessPlatformCompanyAccountGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessPlatformCompanyAccountGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedBusinessPlatformCompanyAccountGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to business_platform_company_account_get_v3.0_account_type value
func (v BusinessPlatformCompanyAccountGetV30AccountType) Ptr() *BusinessPlatformCompanyAccountGetV30AccountType {
	return &v
}
