/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpDataSourceReadV2DataDataListDataSourceType
type DmpDataSourceReadV2DataDataListDataSourceType string

// List of dmp_data_source_read_v2_data_data_list_data_source_type
const (
	UID_DmpDataSourceReadV2DataDataListDataSourceType DmpDataSourceReadV2DataDataListDataSourceType = "UID"
	DID_DmpDataSourceReadV2DataDataListDataSourceType DmpDataSourceReadV2DataDataListDataSourceType = "DID"
)

// All allowed values of DmpDataSourceReadV2DataDataListDataSourceType enum
var AllowedDmpDataSourceReadV2DataDataListDataSourceTypeEnumValues = []DmpDataSourceReadV2DataDataListDataSourceType{
	"UID",
	"DID",
}

// NewDmpDataSourceReadV2DataDataListDataSourceTypeFromValue returns a pointer to a valid DmpDataSourceReadV2DataDataListDataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceReadV2DataDataListDataSourceTypeFromValue(v string) (*DmpDataSourceReadV2DataDataListDataSourceType, error) {
	ev := DmpDataSourceReadV2DataDataListDataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceReadV2DataDataListDataSourceType: valid values are %v", v, AllowedDmpDataSourceReadV2DataDataListDataSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceReadV2DataDataListDataSourceType) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceReadV2DataDataListDataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_read_v2_data_data_list_data_source_type value
func (v DmpDataSourceReadV2DataDataListDataSourceType) Ptr() *DmpDataSourceReadV2DataDataListDataSourceType {
	return &v
}
