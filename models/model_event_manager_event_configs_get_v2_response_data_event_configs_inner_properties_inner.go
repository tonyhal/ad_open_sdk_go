/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerEventConfigsGetV2ResponseDataEventConfigsInnerPropertiesInner struct for EventManagerEventConfigsGetV2ResponseDataEventConfigsInnerPropertiesInner
type EventManagerEventConfigsGetV2ResponseDataEventConfigsInnerPropertiesInner struct {
	// 附加属性描述
	Description *string `json:"description,omitempty"`
	// 附加属性枚举值
	EnumValue *map[string]string `json:"enum_value,omitempty"`
	// 附加属性英文名称
	Field *string `json:"field,omitempty"`
	// 附加属性中文名称
	FieldName *string `json:"field_name,omitempty"`
	// 附加属性单位
	Unit         *string                                                              `json:"unit,omitempty"`
	VariableType *EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType `json:"variable_type,omitempty"`
}
