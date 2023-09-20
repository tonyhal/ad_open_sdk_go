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

// DmpDataSourceUpdateV2OperationType
type DmpDataSourceUpdateV2OperationType int64

// List of dmp_data_source_update_v2_operation_type
const (
	Enum_1_DmpDataSourceUpdateV2OperationType DmpDataSourceUpdateV2OperationType = 1
	Enum_2_DmpDataSourceUpdateV2OperationType DmpDataSourceUpdateV2OperationType = 2
	Enum_3_DmpDataSourceUpdateV2OperationType DmpDataSourceUpdateV2OperationType = 3
)

// All allowed values of DmpDataSourceUpdateV2OperationType enum
var AllowedDmpDataSourceUpdateV2OperationTypeEnumValues = []DmpDataSourceUpdateV2OperationType{
	1,
	2,
	3,
}

// NewDmpDataSourceUpdateV2OperationTypeFromValue returns a pointer to a valid DmpDataSourceUpdateV2OperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceUpdateV2OperationTypeFromValue(v int64) (*DmpDataSourceUpdateV2OperationType, error) {
	ev := DmpDataSourceUpdateV2OperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceUpdateV2OperationType: valid values are %v", v, AllowedDmpDataSourceUpdateV2OperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceUpdateV2OperationType) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceUpdateV2OperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_update_v2_operation_type value
func (v DmpDataSourceUpdateV2OperationType) Ptr() *DmpDataSourceUpdateV2OperationType {
	return &v
}
