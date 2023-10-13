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

// DmpDataSourceCreateV2DataSourceType
type DmpDataSourceCreateV2DataSourceType string

// List of dmp_data_source_create_v2_data_source_type
const (
	UID_DmpDataSourceCreateV2DataSourceType DmpDataSourceCreateV2DataSourceType = "UID"
	DID_DmpDataSourceCreateV2DataSourceType DmpDataSourceCreateV2DataSourceType = "DID"
)

// All allowed values of DmpDataSourceCreateV2DataSourceType enum
var AllowedDmpDataSourceCreateV2DataSourceTypeEnumValues = []DmpDataSourceCreateV2DataSourceType{
	"UID",
	"DID",
}

// NewDmpDataSourceCreateV2DataSourceTypeFromValue returns a pointer to a valid DmpDataSourceCreateV2DataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceCreateV2DataSourceTypeFromValue(v string) (*DmpDataSourceCreateV2DataSourceType, error) {
	ev := DmpDataSourceCreateV2DataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceCreateV2DataSourceType: valid values are %v", v, AllowedDmpDataSourceCreateV2DataSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceCreateV2DataSourceType) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceCreateV2DataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_create_v2_data_source_type value
func (v DmpDataSourceCreateV2DataSourceType) Ptr() *DmpDataSourceCreateV2DataSourceType {
	return &v
}
