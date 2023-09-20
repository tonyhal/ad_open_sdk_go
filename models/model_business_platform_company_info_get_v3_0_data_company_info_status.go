/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus
type BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus string

// List of business_platform_company_info_get_v3.0_data_company_info_status
const (
	EXPIRED_BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus     BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus = "EXPIRED"
	FAILED_BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus      BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus = "FAILED"
	NOT_STARTED_BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus = "NOT_STARTED"
	PROCESSING_BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus  BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus = "PROCESSING"
	SUCCESS_BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus     BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus = "SUCCESS"
	WAITING_BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus     BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus = "WAITING"
)

// All allowed values of BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus enum
var AllowedBusinessPlatformCompanyInfoGetV30DataCompanyInfoStatusEnumValues = []BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus{
	"EXPIRED",
	"FAILED",
	"NOT_STARTED",
	"PROCESSING",
	"SUCCESS",
	"WAITING",
}

// NewBusinessPlatformCompanyInfoGetV30DataCompanyInfoStatusFromValue returns a pointer to a valid BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessPlatformCompanyInfoGetV30DataCompanyInfoStatusFromValue(v string) (*BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus, error) {
	ev := BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus: valid values are %v", v, AllowedBusinessPlatformCompanyInfoGetV30DataCompanyInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus) IsValid() bool {
	for _, existing := range AllowedBusinessPlatformCompanyInfoGetV30DataCompanyInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to business_platform_company_info_get_v3.0_data_company_info_status value
func (v BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus) Ptr() *BusinessPlatformCompanyInfoGetV30DataCompanyInfoStatus {
	return &v
}
