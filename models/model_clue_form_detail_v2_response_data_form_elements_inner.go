/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2ResponseDataFormElementsInner struct for ClueFormDetailV2ResponseDataFormElementsInner
type ClueFormDetailV2ResponseDataFormElementsInner struct {
	AllowEmpty *ClueFormDetailV2DataFormElementsAllowEmpty `json:"allow_empty,omitempty"`
	//
	DefaultValue *int64 `json:"default_value,omitempty"`
	//
	ElementId   *int64                                       `json:"element_id,omitempty"`
	ElementType *ClueFormDetailV2DataFormElementsElementType `json:"element_type,omitempty"`
	IsUnique    *ClueFormDetailV2DataFormElementsIsUnique    `json:"is_unique,omitempty"`
	//
	Label *string `json:"label,omitempty"`
	//
	Layer *int64 `json:"layer,omitempty"`
	//
	Sequence *int64 `json:"sequence,omitempty"`
}