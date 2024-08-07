/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreateStatementV2TemplateType
type CreateStatementV2TemplateType string

// List of create_statement_v2_template_type
const (
	NOT_STANDARD_CreateStatementV2TemplateType CreateStatementV2TemplateType = "NOT_STANDARD"
	STANDARD_CreateStatementV2TemplateType     CreateStatementV2TemplateType = "STANDARD"
)

// All allowed values of CreateStatementV2TemplateType enum
var AllowedCreateStatementV2TemplateTypeEnumValues = []CreateStatementV2TemplateType{
	"NOT_STANDARD",
	"STANDARD",
}

// NewCreateStatementV2TemplateTypeFromValue returns a pointer to a valid CreateStatementV2TemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateStatementV2TemplateTypeFromValue(v string) (*CreateStatementV2TemplateType, error) {
	ev := CreateStatementV2TemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateStatementV2TemplateType: valid values are %v", v, AllowedCreateStatementV2TemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateStatementV2TemplateType) IsValid() bool {
	for _, existing := range AllowedCreateStatementV2TemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to create_statement_v2_template_type value
func (v CreateStatementV2TemplateType) Ptr() *CreateStatementV2TemplateType {
	return &v
}