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

// CreateStatementV2DataResult
type CreateStatementV2DataResult string

// List of create_statement_v2_data_result
const (
	SUCCESS_CreateStatementV2DataResult CreateStatementV2DataResult = "SUCCESS"
	FAILED_CreateStatementV2DataResult  CreateStatementV2DataResult = "FAILED"
)

// All allowed values of CreateStatementV2DataResult enum
var AllowedCreateStatementV2DataResultEnumValues = []CreateStatementV2DataResult{
	"SUCCESS",
	"FAILED",
}

// NewCreateStatementV2DataResultFromValue returns a pointer to a valid CreateStatementV2DataResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateStatementV2DataResultFromValue(v string) (*CreateStatementV2DataResult, error) {
	ev := CreateStatementV2DataResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateStatementV2DataResult: valid values are %v", v, AllowedCreateStatementV2DataResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateStatementV2DataResult) IsValid() bool {
	for _, existing := range AllowedCreateStatementV2DataResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to create_statement_v2_data_result value
func (v CreateStatementV2DataResult) Ptr() *CreateStatementV2DataResult {
	return &v
}
