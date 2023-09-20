/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2ResponseDataForm
type ClueFormDetailV2ResponseDataForm struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Elements    []*ClueFormDetailV2ResponseDataFormElementsInner `json:"elements,omitempty"`
	EnableLayer *ClueFormDetailV2DataFormEnableLayer             `json:"enable_layer,omitempty"`
	ExtendInfo  *ClueFormDetailV2ResponseDataFormExtendInfo      `json:"extend_info,omitempty"`
	FormType    *ClueFormDetailV2DataFormFormType                `json:"form_type,omitempty"`
	//
	InstanceId *int64                         `json:"instance_id,omitempty"`
	IsDel      *ClueFormDetailV2DataFormIsDel `json:"is_del,omitempty"`
	//
	LayerSubmitText *string `json:"layer_submit_text,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SubmitText *string `json:"submit_text,omitempty"`
	//
	UpdateTime   *string                               `json:"update_time,omitempty"`
	ValidateType *ClueFormDetailV2DataFormValidateType `json:"validate_type,omitempty"`
}
