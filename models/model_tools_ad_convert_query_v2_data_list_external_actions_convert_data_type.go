/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertQueryV2DataListExternalActionsConvertDataType
type ToolsAdConvertQueryV2DataListExternalActionsConvertDataType string

// List of tools_ad_convert_query_v2_data_list_external_actions_convert_data_type
const (
	AD_CONVERT_DATA_TYPE_LESS_THAN_TEN_ToolsAdConvertQueryV2DataListExternalActionsConvertDataType     ToolsAdConvertQueryV2DataListExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_LESS_THAN_TEN"
	AD_CONVERT_DATA_TYPE_EVERY_ONE_ToolsAdConvertQueryV2DataListExternalActionsConvertDataType         ToolsAdConvertQueryV2DataListExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_EVERY_ONE"
	AD_CONVERT_DATA_TYPE_ONLY_ONE_ToolsAdConvertQueryV2DataListExternalActionsConvertDataType          ToolsAdConvertQueryV2DataListExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_ONLY_ONE"
	AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW_ToolsAdConvertQueryV2DataListExternalActionsConvertDataType     ToolsAdConvertQueryV2DataListExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW"
	AD_CONVERT_DATA_TYPE_PER_TIMES_ToolsAdConvertQueryV2DataListExternalActionsConvertDataType         ToolsAdConvertQueryV2DataListExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_PER_TIMES"
	AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED_ToolsAdConvertQueryV2DataListExternalActionsConvertDataType ToolsAdConvertQueryV2DataListExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED"
)

// All allowed values of ToolsAdConvertQueryV2DataListExternalActionsConvertDataType enum
var AllowedToolsAdConvertQueryV2DataListExternalActionsConvertDataTypeEnumValues = []ToolsAdConvertQueryV2DataListExternalActionsConvertDataType{
	"AD_CONVERT_DATA_TYPE_LESS_THAN_TEN",
	"AD_CONVERT_DATA_TYPE_EVERY_ONE",
	"AD_CONVERT_DATA_TYPE_ONLY_ONE",
	"AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW",
	"AD_CONVERT_DATA_TYPE_PER_TIMES",
	"AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED",
}

// NewToolsAdConvertQueryV2DataListExternalActionsConvertDataTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2DataListExternalActionsConvertDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2DataListExternalActionsConvertDataTypeFromValue(v string) (*ToolsAdConvertQueryV2DataListExternalActionsConvertDataType, error) {
	ev := ToolsAdConvertQueryV2DataListExternalActionsConvertDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2DataListExternalActionsConvertDataType: valid values are %v", v, AllowedToolsAdConvertQueryV2DataListExternalActionsConvertDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2DataListExternalActionsConvertDataType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2DataListExternalActionsConvertDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_data_list_external_actions_convert_data_type value
func (v ToolsAdConvertQueryV2DataListExternalActionsConvertDataType) Ptr() *ToolsAdConvertQueryV2DataListExternalActionsConvertDataType {
	return &v
}
